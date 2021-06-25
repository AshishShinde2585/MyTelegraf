package test

import "fmt"
import "github.com/influxdata/telegraf"

type serializer struct {
	Prefix string
}

//only return structure which implmemets two methods
func NewSerializer(prefix string) (*serializer, error) {
	s := &serializer{
		Prefix : prefix,
	}
	return s, nil
}

func (s *serializer) Serialize(metric telegraf.Metric) ([]byte, error) {

	output_string := "\n{"+ s.Prefix + ": \n Measurement => "+ metric.Name()
	output_string = output_string + "\n Tags => "
	for tag,val := range metric.Tags() {
		output_string = output_string+ "\n\t"+ tag + "=" + val
	}
	output_string = output_string + "\n Fields => "
	for tag,val := range metric.Fields() {
		output_string = output_string+ "\n\t"+ tag + "=" + fmt.Sprintf("%v", val)
	}
	output_string = output_string + "\n}"
	serialized := []byte(output_string)

	return serialized, nil
}

func (s *serializer) SerializeBatch(metrics []telegraf.Metric) ([]byte, error) {
	return []byte{}, nil
}