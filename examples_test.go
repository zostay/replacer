package replacer_test

import (
	"fmt"
	"strings"

	"github.com/zostay/replacer"
)

func Example_bad() {
	// DO NOT DO THIS: This will give you inconsistent results.
	r := strings.NewReplacer(
		"modifiedbyid", "modified_by_id",
		"modifiedby", "modified_by")
	fmt.Println(r.Replace("modifiedbyid"))
	// Output may be: modified_by_id
	//            OR: modified_byid
}

func Example_good() {
	// DO THIS! This will always give you the same result.
	r := replacer.New(
		"modifiedbyid", "modified_by_id",
		"modifiedby", "modified_by")
	fmt.Println(r.Replace("modifiedbyid"))
	// Output will always be: modified_by_id
}
