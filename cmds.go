package main

func createCommands(devices []deviceStruct) []string {
	// TODO PLACEHOLDER
	commands := []string{}

	for _, s := range devices {
		command := "lpadmin" + " -P " + s.name + " -v //:" + s.ip

		if s.ppdNeeded {
			// Add address for PPD file
			if s.ppdType == "file" {

			} else {
				// CONTINUE FROM HERE

			}
		}

		if len(s.options) > 0 {
			// Add options according to their number
		}

		commands = append(commands, command)
	}

	return commands
}
