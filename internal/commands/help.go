package commands

const (
	landOfLifeInformationChart = `
**Land of Life Schedule**
────────────────────────

**1st LoL = 01:00 AM**
Start: <t:64060581600:t>
Asgobas: <t:64060585200:t>
End: <t:64060588800:t>

**2nd LoL = 03:00 AM**
Start: <t:64060588800:t>
Asgobas: <t:64060592400:t>
End: <t:64060596000:t>

**3rd LoL = 05:00 AM**
Start: <t:64060596000:t>
Asgobas: <t:64060599600:t>
End: <t:64060603200:t>

**4th LoL = 07:00 AM**
Start: <t:64060603200:t>
Asgobas: <t:64060606800:t>
End: <t:64060606800:t>

**5th LoL = 09:00 AM**
Start: <t:64060606800:t>
Asgobas: <t:64060614000:t>
End: <t:64060617600:t>

**6th LoL = 11:00 AM**
Start: <t:64060617600:t>
Asgobas: <t:64060621200:t>
End: <t:64060624800:t>

**7th LoL = 13:00 PM**
Start: <t:64060624800:t>
Asgobas: <t:64060628400:t>
End: <t:64060632000:t>

**8th LoL = 15:00 PM**
Start: <t:64060632000:t>
Asgobas: <t:64060635600:t>
End: <t:64060639200:t>

**9th LoL = 17:00 PM**
Start: <t:64060639200:t>
Asgobas: <t:64060642800:t>
End: <t:64060646400:t>

**10th LoL = 19:00 PM**
Start: <t:64060646400:t>
Asgobas: <t:64060650000:t>
End: <t:64060653600:t>

**11th LoL = 21:00 PM**
Start: <t:64060653600:t>
Asgobas: <t:64060657200:t>
End: <t:64060660800:t>

**12th LoL = 23:00 PM**
Start: <t:64060660800:t>
Asgobas: <t:64060664400:t>
End: <t:64060668000:t>

────────────────────────
Times are displayed in your local timezone
Board might look broken for you, it's optimized for 24h format (00:00)
LoL is available on all channels
`
	statusCommandInstructions     = "\n\nTo view the status of today's slots, use the `/lol slot status` command. You can optionally filter by hour to see specific time slots."
	registerCommandInstructions   = "\n\nTo register for a slot, use the `/lol slot register` command with your in-game username, desired hour (e.g., 01:00, 03:00, or just 1,3,5), channel number (1-7), character class, level, and pet name."
	unregisterCommandInstructions = "\n\nTo unregister from a slot, use the `/lol slot unregister` command with your in-game username, hour, and channel number."
)
