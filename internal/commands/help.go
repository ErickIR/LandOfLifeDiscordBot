package commands

const (
	landOfLifeInformationChart = `
**Land of Life Schedule**
────────────────────────

**1st LoL = 01:00 AM CEST/CET**
Start: <t:64060581600:t>
Asgobas: <t:64060585200:t>
End: <t:64060588800:t>

**2nd LoL = 03:00 AM CEST/CET**
Start: <t:64060588800:t>
Asgobas: <t:64060527600:t>
End: <t:64060531200:t>

**3rd LoL = 05:00 AM CEST/CET**
Start: <t:64060531200:t>
Asgobas: <t:64060534800:t>
End: <t:64060538400:t>

**4th LoL = 07:00 AM CEST/CET**
Start: <t:64060538400:t>
Asgobas: <t:64060542000:t>
End: <t:64060545600:t>

**5th LoL = 09:00 AM CEST/CET**
Start: <t:64060545600:t>
Asgobas: <t:64060549200:t>
End: <t:64060552800:t>

**6th LoL = 11:00 AM CEST/CET**
Start: <t:64060552800:t>
Asgobas: <t:64060556400:t>
End: <t:64060560000:t>

**7th LoL = 13:00 PM CEST/CET**
Start: <t:64060560000:t>
Asgobas: <t:64060563600:t>
End: <t:64060567200:t>

**8th LoL = 15:00 PM CEST/CET**
Start: <t:64060567200:t>
Asgobas: <t:64060570800:t>
End: <t:64060570800:t>

**9th LoL = 17:00 PM CEST/CET**
Start: <t:64060570800:t>
Asgobas: <t:64060578000:t>
End: <t:64060581600:t>

**10th LoL = 19:00 PM CEST/CET**
Start: <t:64060581600:t>
Asgobas: <t:64060585200:t>
End: <t:64060588800:t>

**11th LoL = 21:00 PM CEST/CET**
Start: <t:64060588800:t>
Asgobas: <t:64060592400:t>
End: <t:64060596000:t>

**12th LoL = 23:00 PM CEST/CET**
Start: <t:64060596000:t>
Asgobas: <t:64060599600:t>
End: <t:64060516800:t>

────────────────────────
Times are displayed in your local timezone
Board might look broken for you, it's optimized for 24h format (00:00)
LoL is available on all channels
`
	statusCommandInstructions     = "\n\nTo view the status of today's slots, use the `/lol slot status` command. You can optionally filter by hour to see specific time slots."
	registerCommandInstructions   = "\n\nTo register for a slot, use the `/lol slot register` command with your in-game username, desired hour (e.g., 01:00, 03:00, or just 1,3,5), channel number (1-7), character class, level, and pet name."
	unregisterCommandInstructions = "\n\nTo unregister from a slot, use the `/lol slot unregister` command with your in-game username, hour, and channel number."
)
