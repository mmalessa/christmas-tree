package christmastree

type ChristmasTreePattern struct {
	template string
	config   map[string]interface{}
}

func (ch *ChristmasTree) AddPattern(patternid string, templatename string, config map[string]interface{}) error {
	ch.patterns[patternid] = ChristmasTreePattern{
		template: templatename,
		config:   config,
	}
	return nil
}

func (ch *ChristmasTree) PlayPattern(patternid string) error {
	pattern := ch.patterns[patternid]
	template := pattern.template

	switch template {
	case "wipe":
		return ch.PlayTemplateWipe(pattern.config)
	case "vwipe":
		return ch.PlayTemplateVWipe(pattern.config)
	case "rainbow":
		return ch.PlayTemplateRainbow(pattern.config)
	case "sleep":
		return ch.PlayTemplateSleep(pattern.config)
	case "ruler":
		return ch.PlayTemplateRuler(pattern.config)
	case "rainbowunicorn":
		return ch.PlayTemplateRainbowUnicorn(pattern.config)
	}

	return nil
}
