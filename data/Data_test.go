package Data

import (
	"fmt"
	"testing"
)

func Test_GenSourceData(t *testing.T) {

	pmnt := GenPaymentSource()

	fmt.Printf("%+v", pmnt)

	// s := GetSpot("Rest Bay (Porthcawl)")

	// if s.Name != "Rest Bay (Porthcawl)" {
	// 	t.Errorf("invalid name : %s", s.Name)
	// }

}
