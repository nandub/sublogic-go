package sublogic

import (
    "math"
    "errors"
)

type Opinion struct {
	Belief      float64
	Disbelief   float64
	Uncertainty float64
	Baserate    float64
}

// Creates a new opinion.
// 
// Opinion must satisfy the following:
// - belief + disbelief + uncertainty = 1
// - Each component must be in the range of 0 to 1
func NewOpinion(belief, disbelief, uncertainty, baserate float64) (*Opinion, error) {
	// Test for valid opinion
	if val := belief + disbelief + uncertainty; val != 1 {
		return newEmptyOpinion(), errors.New("The sum of all opinion components doesn't equal 1")
	}
	
	if outOfRange(belief) || outOfRange(disbelief) || outOfRange(uncertainty) || outOfRange(baserate) {
		return newEmptyOpinion(), errors.New("Opinion component isn't in range of 0 to 1")
	}
	
	opinion := Opinion{ belief, disbelief, uncertainty, baserate }
	return &opinion, nil
}

func newEmptyOpinion() (*Opinion) {
	return &Opinion{ }
}

// Where A trusts B, and B trusts X, this returns A's transitive trust in X
func (A *Opinion) Discount(B *Opinion) (C *Opinion) {
	C = newEmptyOpinion()
	
	C.Belief = A.Belief * B.Belief
	C.Disbelief = A.Belief * B.Disbelief
	C.Uncertainty = (A.Disbelief + A.Uncertainty + A.Belief*B.Uncertainty)
	C.Baserate = B.Baserate
    
    C.CheckConsistency()
	
	return C
}

func (this *Opinion) CheckConsistency() {
    this.Belief = constrain(adjust(this.Belief))
    this.Disbelief = constrain(adjust(this.Disbelief))
    this.Uncertainty = constrain(adjust(this.Uncertainty))
    if math.Abs(float64(this.Belief + this.Disbelief + this.Uncertainty - 1.0)) > 1.0E-010 {
        var bdu float64
        bdu = this.Belief + this.Disbelief + this.Uncertainty
        this.Belief = constrain(adjust(this.Belief / bdu))
        this.Uncertainty = constrain(adjust(this.Uncertainty / bdu))
        this.Disbelief = (1.0 - (this.Belief + this.Uncertainty))
    }
}
