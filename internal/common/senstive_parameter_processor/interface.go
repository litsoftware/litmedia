package senstive_parameter_processor

type Encoder interface {
	Encode(params map[string]string, data interface{}) error
}

type Decoder interface {
	Decode(params map[string]string, data interface{}) error
}

type Processor interface {
	Encoder
	Decoder
}
