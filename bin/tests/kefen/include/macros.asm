; Macro to use serial registers for printing out data.
MACRO OutputA
	PUSH AF
	LD [$ff00+$01], A
	LD A, $80
	LD [$ff00+$02], A
	POP AF
ENDM