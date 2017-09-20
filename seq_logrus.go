package logrus_seq_hook

import (
	"github.com/mundipagg/goseq"
	"github.com/sirupsen/logrus"
)

//SeqHook is a struct to support seq logrus hook
type SeqHook struct {
	writer *goseq.Logger
	binds  map[logrus.Level]func(*logrus.Entry, *goseq.Logger) error
}

//NewSeqHook creates a new hook from url and api token
func NewSeqHook(seqURL string, apiKey string) (hook *SeqHook, err error) {
	hook = new(SeqHook)
	hook.writer, err = goseq.GetLogger(seqURL, apiKey)
	hook.binds = make(map[logrus.Level]func(*logrus.Entry, *goseq.Logger) error)

	hook.binds[logrus.DebugLevel] = seqDebug
	hook.binds[logrus.InfoLevel] = seqInfo
	hook.binds[logrus.WarnLevel] = seqWarn
	hook.binds[logrus.FatalLevel] = seqFatal

	return
}

//Levels return available levels on SEQ
func (seq *SeqHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.FatalLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

func (seq *SeqHook) Fire(entry *logrus.Entry) error {
	return seq.binds[entry.Level](entry, seq.writer)
}

func convertFieldToPropertie(entry *logrus.Entry) goseq.Properties {
	props := goseq.NewProperties()
	for k, v := range entry.Data {
		props.AddProperty(k, v)
	}
	return props
}

func seqDebug(entry *logrus.Entry, seq *goseq.Logger) error {
	seq.Debug(entry.Message, convertFieldToPropertie(entry))
	return nil
}
func seqInfo(entry *logrus.Entry, seq *goseq.Logger) error {
	seq.Information(entry.Message, convertFieldToPropertie(entry))
	return nil
}
func seqWarn(entry *logrus.Entry, seq *goseq.Logger) error {
	seq.Warning(entry.Message, convertFieldToPropertie(entry))
	return nil
}

func seqFatal(entry *logrus.Entry, seq *goseq.Logger) error {
	seq.Fatal(entry.Message, convertFieldToPropertie(entry))
	return nil
}
