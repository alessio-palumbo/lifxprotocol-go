// Code generated. DO NOT EDIT.
// Source: https://github.com/LIFX/public-protocol@0000000000000000000000000000000000000000
// Generated: 2006-01-02T15:04:05Z
package enums

type TestEnum uint8

const (
	TestEnumFirstValue  TestEnum = 1
	TestEnumSecondValue TestEnum = 2
	TestEnumThirdValue  TestEnum = 3
)

func (e TestEnum) String() string {
	switch e {
	case TestEnumFirstValue:
		return "FirstValue"
	case TestEnumSecondValue:
		return "SecondValue"
	case TestEnumThirdValue:
		return "ThirdValue"
	}
	return ""
}
