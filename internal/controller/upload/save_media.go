package upload

import (
	"context"
	"github.com/google/uuid"
	"github.com/litsoftware/litmedia/internal/common/types"
	"github.com/litsoftware/litmedia/internal/controller"
	"github.com/litsoftware/litmedia/internal/errors"
	"github.com/litsoftware/litmedia/internal/g"
	"github.com/litsoftware/litmedia/internal/pkg/config/orm"
	"io"
	"os"
)

type saveMedia struct {
	controller.BaseProc
}

func (c *saveMedia) Process() {
	r := <-c.In
	st := r.Req.(*Statement)
	if st.Err != nil {
		c.Next(r)
		return
	}

	var err error
	ctx, cancel := context.WithCancel(r)
	defer cancel()

	var IsEncrypted = false
	var level = types.MediaLevel(st.Request.Level)
	if level == types.MediaLevel_Sensitive {
		IsEncrypted = true
	}

	us := NewUploadService()
	tmpFileName := st.fileHeader.Filename + ".tmp"
	tmp, err := os.Create(tmpFileName)
	if err != nil {
		g.App.Logger.Errorf("create tmp file error: %s", err.Error())
		st.Err = errors.ErrorInternalError(err.Error())
		c.Next(r)
		return
	}
	_, err = io.Copy(tmp, st.file)
	if err != nil {
		g.App.Logger.Errorf("create tmp file error: %s", err.Error())
		st.Err = errors.ErrorInternalError(err.Error())
		c.Next(r)
		return
	}
	defer func() {
		_ = tmp.Close()
		_ = os.Remove(tmpFileName)
	}()

	err = us.Parse(tmp, level)
	if err != nil {
		st.Err = err
		c.Next(r)
		return
	}

	st.Media, err = orm.GetClient().Media.
		Create().
		SetAppID(st.App.ID).
		SetUserID(st.App.OperatorID).
		SetFileName(us.GetFileName()).
		SetSize(us.GetSize()).
		SetLevel(int(level.Number())).
		SetIsEncrypted(IsEncrypted).
		SetSavePath(us.GetSavePath()).
		SetExt(us.GetExt()).
		SetFullPath(us.GetFullPath()).
		SetHash(us.GetCheckSum()).
		SetMime(us.GetMime().String()).
		SetOrgFileName(st.Request.Filename).
		SetType(int(us.GetMediaType().Number())).
		SetStatus(1).
		SetRefCount(1).
		SetReason("").
		SetSn(uuid.New().String()).
		Save(ctx)

	if err != nil {
		st.Err = err
		c.Next(r)
		return
	}

	c.Next(r)
	return
}
