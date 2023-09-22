package cruelerror

type CruelError int

func (CruelError) Error() string {
	return "are you an idiot?!"
}
