package display

type DisplayOptions struct {
	args OptionsArgs
}

type OptionsArgs struct {
	Words bool
	Lines bool
	Bytes bool
}

func (this DisplayOptions) ShowWords() bool {
	return this.args.Words
}

func (this DisplayOptions) ShowLines() bool {
	return this.args.Lines
}

func (this DisplayOptions) ShowBytes() bool {
	return this.args.Bytes
}

func (this DisplayOptions) ShowAll() bool {
	return this.ShowWords() || this.ShowLines() || this.ShowBytes()
}

func NewOptionArgs() OptionsArgs {
	return OptionsArgs{Words: true, Lines: true, Bytes: true}
}

func NewOptions(args OptionsArgs) DisplayOptions {
	return DisplayOptions{
		args: args,
	}
}

func (opts DisplayOptions) WithDefaults() DisplayOptions {
	if !opts.args.Words && !opts.args.Lines && !opts.args.Bytes {
		return DisplayOptions{args: OptionsArgs{Words: true, Lines: true, Bytes: true}}
	}
	return opts
}
