package unit

type Fraction struct {
	Num int
	Den int
}

func (f Fraction) Add(other Fraction) Fraction {
	result := Fraction{
		Num: f.Num*other.Den + other.Num*f.Den,
		Den: f.Den * other.Den,
	}

	return result.Reduce()
}

func (f Fraction) Reduce() Fraction {
	if f.Den == 0 {
		return f
	}
	if f.Num == 0 {
		return Fraction{Num: 0, Den: 1}
	}

	g := GCD(f.Num, f.Den)
	reduced := Fraction{
		Num: f.Num / g,
		Den: f.Den / g,
	}

	if reduced.Den < 0 {
		reduced.Num = -reduced.Num
		reduced.Den = -reduced.Den
	}

	return reduced
}
