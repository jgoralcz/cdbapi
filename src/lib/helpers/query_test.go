package helpers

import (
	"errors"
	"testing"
)

func TestMaxLimitEmptyString(t *testing.T) {
	queryLimit := ""
	expected := 1

	emptyLimit := MaxLimit(queryLimit, 1, 20)

	if emptyLimit != expected {
		t.Errorf("MaxLimit(\"\", 1, 20) failed. Expected %v, Received %v", expected, emptyLimit)
	} else {
		t.Logf("MaxLimit(\"\", 1, 20) passed.")
	}
}

func TestMaxLimitBadString(t *testing.T) {
	queryLimit := "test"
	expected := 1

	limit := MaxLimit(queryLimit, 1, 20)

	if limit != expected {
		t.Errorf("MaxLimit(\"test\", 1, 20) failed. Expected %v, Received %v", expected, limit)
	} else {
		t.Logf("MaxLimit(\"test\", 1, 20) passed.")
	}
}

func TestMaxLimitOverMaxString(t *testing.T) {
	queryLimit := "300"
	expected := 20

	limit := MaxLimit(queryLimit, 1, 20)

	if limit != expected {
		t.Errorf("MaxLimit(\"300\", 1, 20) failed. Expected %v, Received %v", expected, limit)
	} else {
		t.Logf("MaxLimit(\"300\", 1, 20) passed.")
	}
}

func TestMaxLimitLessThanString(t *testing.T) {
	queryLimit := "2"
	expected := 5

	limit := MaxLimit(queryLimit, 5, 20)

	if limit != expected {
		t.Errorf("MaxLimit(\"2\", 5, 20) failed. Expected %v, Received %v", expected, limit)
	} else {
		t.Logf("MaxLimit(\"2\", 5, 20) passed.")
	}
}

func TestMaxLimitBigString(t *testing.T) {
	queryLimit := "123456789"
	defaultLimit := 5
	maxLimit := 20
	expected := 5

	received := MaxLimit(queryLimit, defaultLimit, maxLimit)

	if received != expected {
		t.Errorf("MaxLimit(%v, %v, %v) failed. Expected %v, Received %v", expected, defaultLimit, maxLimit, expected, received)
	} else {
		t.Logf("MaxLimit(%v, %v, %v) passed.", queryLimit, defaultLimit, maxLimit)
	}
}

func BenchmarkMaxLimit4String(b *testing.B) {
	str4 := "1234"
	for i := 0; i < b.N; i++ {
		MaxLimit(str4, 1, 20)
	}
}

func BenchmarkMaxLimit1024String(b *testing.B) {
	str1024 := "RivB8d4iEuol9MLFllafqqAmdYqfXcUY5NzWhBrHabxP4RYceelq8F5NxYbSr56OyCOWgYfmquYsQSPyNyY5OXMjDHwmqDCocEKl8auiznqbZexviKvXDap62DuWemjQqpUpAntEzBahzulQmzYj01E1XCoUJvanvT4abnUktXr4XXUk4tzeOI78WdrD6FVUvN1eyMppfvRK1ScF0VKeCCIFLYEaG8b4BQYuU1HlcSuyE4fNSDpdwePwgYl3Fvb811uRhr4Fo374fArBTOc5YEyHMEO95Vo1dCjtItmPhFLdGcgrAENrPMZ5vh5J5llabH4dqE1N0Ej090w1POoGcUWeUsmYQ8enGSryLjcyuKLeAeLWi3SblLm4pBt5UnVaiHxNiZslfpLLPFKGD5j0DcrTCzX50DU6SumXXk8ON9K9KVy2q9GbrUvjvxi8V687mP37bax7UpV0j2dHkZ5XlG01GELPOpcEDoaltZ8XdzjXRjuRs1JuyaqoFIhp6bCP79bCOnCgBrNXxGzy0rgHo2USXD98NVlQW1v9H2zabVJK0UljMolXwAlBz7GRR4lQwuhge6OtujlJQK1Icg5mHpYyPaTk0Lj0luy0YwLbuwnbBjKGKneP3zLiwzuOyNpVDWwq8uuYhusQzS6arOFJxJTGbhOCNchTrDUbuwpvIuq0jXt8aIyXa1N9FiZeBegT3AVcpneMWNgC8fcvkaWG2Lo8DtqhCvp8hqSh6VcQDRa7XvG6pMcmRDXgkiyEKFt4Gatmiu7OcdNpEBOBTuFQ4EIieFImO5LIWeqiz9UWMQCMiHiKcOAoDXWasTPvh5j7aUhDbkr6McyeVUobxXixckzo1IY24SK2gBlEJvRbTwqImd4zAjgdbRaOL4vhTZjdmf4elvt5mq0gvI5pM8GULL3r4eKHLTJ1f9HpvUEePI4Dz8le4dgbWLnRwTLaxy0cxew0R18zapOLRWJCeSSuXUuQNRmIP2VXGGZrmtGBmTrGIf1UbvYxUw7GpamKt0W0"
	for i := 0; i < b.N; i++ {
		MaxLimit(str1024, 1, 20)
	}
}

func BenchmarkMaxLimit4096String(b *testing.B) {
	str4096 := "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"
	for i := 0; i < b.N; i++ {
		MaxLimit(str4096, 1, 20)
	}
}

//// DefaultBoolean

func TestDefaultBooleanTrueString(t *testing.T) {
	expected := "true"

	boolStr := DefaultBoolean("true")

	if boolStr != expected {
		t.Errorf("DefaultBoolean(\"true\") failed. Expected %v, Received %v", expected, boolStr)
	} else {
		t.Logf("DefaultBoolean(\"true\") passed.")
	}
}

func TestDefaultBooleanFalseString(t *testing.T) {
	expected := "false"

	boolStr := DefaultBoolean("false")

	if boolStr != expected {
		t.Errorf("DefaultBoolean(\"false\") failed. Expected %v, Received %v", expected, boolStr)
	} else {
		t.Logf("DefaultBoolean(\"false\") passed.")
	}
}

func TestDefaultBooleanFalseUpperCaseString(t *testing.T) {
	expected := "false"

	boolStr := DefaultBoolean("FaLsE")

	if boolStr != expected {
		t.Errorf("DefaultBoolean(\"FaLsE\") failed. Expected %v, Received %v", expected, boolStr)
	} else {
		t.Logf("DefaultBoolean(\"FaLsE\") passed.")
	}
}

func TestDefaultBooleanAnyString(t *testing.T) {
	expected := ""

	boolStr := DefaultBoolean("test")

	if boolStr != expected {
		t.Errorf("DefaultBoolean(\"test\") failed. Expected %v, Received %v", expected, boolStr)
	} else {
		t.Logf("DefaultBoolean(\"test\") passed.")
	}
}

func BenchmarkDefaultBoolean8String(b *testing.B) {
	str8 := "12345678"
	for i := 0; i < b.N; i++ {
		DefaultBoolean(str8)
	}
}

func BenchmarkDefaultBoolean1024String(b *testing.B) {
	str1024 := "RivB8d4iEuol9MLFllafqqAmdYqfXcUY5NzWhBrHabxP4RYceelq8F5NxYbSr56OyCOWgYfmquYsQSPyNyY5OXMjDHwmqDCocEKl8auiznqbZexviKvXDap62DuWemjQqpUpAntEzBahzulQmzYj01E1XCoUJvanvT4abnUktXr4XXUk4tzeOI78WdrD6FVUvN1eyMppfvRK1ScF0VKeCCIFLYEaG8b4BQYuU1HlcSuyE4fNSDpdwePwgYl3Fvb811uRhr4Fo374fArBTOc5YEyHMEO95Vo1dCjtItmPhFLdGcgrAENrPMZ5vh5J5llabH4dqE1N0Ej090w1POoGcUWeUsmYQ8enGSryLjcyuKLeAeLWi3SblLm4pBt5UnVaiHxNiZslfpLLPFKGD5j0DcrTCzX50DU6SumXXk8ON9K9KVy2q9GbrUvjvxi8V687mP37bax7UpV0j2dHkZ5XlG01GELPOpcEDoaltZ8XdzjXRjuRs1JuyaqoFIhp6bCP79bCOnCgBrNXxGzy0rgHo2USXD98NVlQW1v9H2zabVJK0UljMolXwAlBz7GRR4lQwuhge6OtujlJQK1Icg5mHpYyPaTk0Lj0luy0YwLbuwnbBjKGKneP3zLiwzuOyNpVDWwq8uuYhusQzS6arOFJxJTGbhOCNchTrDUbuwpvIuq0jXt8aIyXa1N9FiZeBegT3AVcpneMWNgC8fcvkaWG2Lo8DtqhCvp8hqSh6VcQDRa7XvG6pMcmRDXgkiyEKFt4Gatmiu7OcdNpEBOBTuFQ4EIieFImO5LIWeqiz9UWMQCMiHiKcOAoDXWasTPvh5j7aUhDbkr6McyeVUobxXixckzo1IY24SK2gBlEJvRbTwqImd4zAjgdbRaOL4vhTZjdmf4elvt5mq0gvI5pM8GULL3r4eKHLTJ1f9HpvUEePI4Dz8le4dgbWLnRwTLaxy0cxew0R18zapOLRWJCeSSuXUuQNRmIP2VXGGZrmtGBmTrGIf1UbvYxUw7GpamKt0W0"
	for i := 0; i < b.N; i++ {
		DefaultBoolean(str1024)
	}
}

func BenchmarkDefaultBoolean4096String(b *testing.B) {
	str4096 := "6ZNvT9Z5uOHRhZCN9XMXyqDLDncfLSOPHj1A3BrVDlQ9Q0QqIa1evqSl2Gds48IDwMiNNB3cXiojstNwAGY0H0W8A1WeRflyROrfeakjotW5XHvrmVMVBMeutuZ13dZUZMUYvZc0UM2QmqaS4RFUP4muUfixkHh3CPll5HZqiOsW84fXv2E9DhmBSyNyga0caJhWPzvra4xLjQjLjzxXGh3ED4fAbpgM9HWASzUp1ivgYdzoOCdR0T39O202u3B18FSub5jUw3TLrnBrnbJbk50hm1xiGYIpyDYcuQpM2j7Uiwf16uI2E9fi1mjCfv4o0zUnuFM7hdeb4CJu2MYNlNZnI2phWwnaKffRMv6ScMonL9RFR549GGPuT8XGxvYLTEGy0m3fTOCjuxQmFSmsQdgRPtmqHnUnhwZfRTHGX8pHblpxG2bNktz5zhOwZCc0hFULzEhvxpOX0Rvv30LJnvJ6GQMMXXJ9ZAPOu10wLhG6k78f8BQR04mBknM4rkTeAjssa5CdcJoPPkFPaINAu3jveGrKsztHARlGl6STHTDJ3Kt96N3qkbM6X6dSRTj8MGnHAKkKR8IiziXVopFegKU8053LaPUo5IYVp6iZ12ryItvIn3MZnkD7lIiVvkQiw6z0raDcA0U5RTqfA0j86dVJSLoAydnbpxgURC6hEZyAeWYMUYknZox3V365HvgrpL3ajPRIvgLkIo7CWLZ92f12fP9dI2wOMkXeQ0p35GHy7Mk3BAqBoodfls8Jxwhs943NDz9INDOdJ0lkyCS205w8ANtcDfLsiG3SXpwL1GcoGRPhhmeybYH4MpgRCgrMB5QfaA98gpiSHuDcOx5Mbl3RySPzgGzqP7moh21JkU36jpCa1A1NHFjDklx9n4imYcgxt0VFNUszM8vxbcDqhWtKz71lkufE0UEsig46V4pyIsFU1XLvx56AbrVLjB3pVcLPTw0uuKyVrX8CQhTuvQF8ST9KLY5ciPPlVmF30Ht0nzYd9tep7OJ3ALgvS1mwaVkTmOx1cd5hbR49oclPqs0JtTbQYfCw9LawlgSrbwdbed3ZHVqwCfqTDDxfkDpcLA6epAK5gHsW3wlFdXgasXfFgqzMp2FNMOdrpMSs7HFd11qeYSlTk68uJDcqZPielMrNMZAcwARIpCXUFVGv0NlDK4mJnZtFVuvbttOpMkCF0c6RIe7dklz8mbm85FUcQeP1zHW7mLsK4fTnRE5EDQHBzwgLi1SM6YOv7wXw3kGz4BBmQj1AguzmAavEjxoXgAbtX7jfPBrYe1kL8s7FiqZNpeOPBapJA4VR6B22v4tRQwpR9RR4lccRwm5KVhzUDnTnSWMxVdfKzYSrpydZgakPXaIv48kFtL8jOjM6UvDftiN65IiPjlIqEMMNxvjueim2jl49RnSUnmmIYRrvFEHaUeKI5rkh4ZmFIgWQvLO1ulOi6LlAP2H79jeVoDKAj0ALRt8myiiM2EXiAiVdkznqAhq55H6cUMrE9BFiglauB0jxJVh2FPI53US1kRi45QE3DVAJ9br7QHTO71MpEalGNI11TfNVFCaNXWds2tazNCl6WphF2uw8TeuUYEeVdm3EYUlaTPrOkEGWnCSBn1lAuJOfmAwkpO3E0GXs1LTDFVVXsjs3TBKGpllvZGmVTUXdv99o8mPcn8JqL0bT4HPsJnXma4XWTzhktOEd7pIGNuNkSHMZ6zmzhSRdWIwgbQR2Gk59HZpXllAi8bNojvnYrR9cR05AtE4i3CKr1GfWjyOrc4nVbHCceYlZqFbStOdmzkqhxBtYHwSrFZ4x6USCkxfwvr8opHlOOgQvWTllXArxGI8qos5aP79zTHSvoIXeY50r9mXyQrrQ4Gr7rJTdLrj6eTOArLUnuKMkyqpKylV7uVDYkyeXsswSauM3woRhXdkLRlqgF9oPeaXIqBP27o12lc8payXMbjGTEt90jDL1I9haqynsfboz0S1UHoVBO0stTs1Gq0s1kPW6AzmsviWDshXEkNDISxqwc7vFuyiYwlRrowAWXKd3gHil5KMU47bKfdBvYVYm6nIl7kTWAQY0aj0qldb1HbONXN4N1XKPRYCnMsYaszLJ7Bp8znAz0ryl7mSFZQTkqrR0yQ2sxI7xcJTmJFL20TT8G78MzSdwO4du61L5utRfXIXqEQTPKowjneLQtiMQnrtCJeEDD3n1FScrFkEWT1jQWFx1EQffmEGD2bzZaBQXFEfTztP9EUNLEKiMHOiJr6ApIyiYUYUDCcqwEvKsLD0t2e7IifE6udhjNwBZ3iej88CrtKJqNLtWupeHQnNpBgZZBuk7uxGoh9VaaxaWo6Qm04JCYnmovPoYTCPCRWShYWr1cAyoPXX6s8AccyZ8FK3j6l7qLUfG0kdhFjfpsQmkK9CFazuzip7NRRxrZU9Ssd4n1XczcVKFdyKsRoHkjE7hxoog61ybL4H7PXOx1BMDkKHHoQayc1lpwoFNZTAYcFqjafRZvxfWvCCuNoF6EKkxaR7sR38HgrvAmCiAUIYcCGNYXdhCxUhDjAl6Nn003ctbnRjEpqNlERAVgeeKZxZjoMlm2YbE1UMx54ggCdaQGExMW1Slbb01XnU5ZSMEgoq6Kw99AEDmzMUntj2jCy4pOfPKuyfhMANVhi8o93dGO5HRcAP9dhy4zLs6rbt0KdrMUBQUS3SDzf3bdL611VAf3e51QAu3d6bIHQYkB1WjIqc4Kz9hMwn4UE7MgfyCTkiB8KXGWxnmh8aZ2faRCoo2nVKIdoQEEGauGc6NXDlUklg43r3lgedVoUt2vGOEWfl91Nndx8BHcOvSsNzt9jNqfvtBNfkJg1K4ihxZ8eiRnuD9TCLmQf1CXj4Y3oapomvbgo5zFQnXR0iyB9KIop7jpBjHnG3bza6lxzRYWmxPGxG58p4pp1kMnIY8uSI0AZsbwN8IsJCjltVRdlXpog7ecLhOwS5hJ76OrE6eHHufXrDg6m8X8tH1PWysEB8Za5JoKR1YR4JqIViUUehY4nPO2MVstOtWbVq3KivG1ajxANTFhtdMAOMtoFX1Kes67mEYJPHraW01KNL59DhQzM7caolxLGrB8GFcIZealrqMVtRAhu2E1brePxw9MqkvNs56b0VEtu7brqAem2iC3jhR4qS8hiAUQ3U7l24BOQsuxontrSIDs4r1t49WGRBc2DnHMNnjYNlxrkrymHR8CCfHRstmqc9hspuoR3wTqyDriE9yk9jUYVWbDbnvI47Ug3mPSD36dPgBiw2wrVeENksNla3Lfm2xfLpp0TCYmz8YnvUmn1cXaPatD9O0ebTWsTO036VmpeGUogGRMjxJnOlePOOq7mSHyDUfkFzINYqtKMNCBd4C00avXPatbAhXSfY7DH56qnHoI9bIBXM5X6ZTSttihoz2ZkQq2oBXQY48LBKs2VwahrL2Ybxc5Q8DWoGgHuZYDMrzj0s9y682BFTPqdti98El3wgClrTkhP6xaoUWJaa8WKMKK121RrBhiPrP8l80gVHWt9SA1JlCNbri6wZVR4tXXyRpvJEVDwdaaR1fQYr0ojd9FBBZ7WVyvw0f89PXQ73HPYHbnOYm4uocM6kpTc6V7N40MuCfoQ9dhti6jmprtuJevLs3W4ODXaEx2KqBxYwabXLMBflsWm36zOBzCT6beKPcv9WZqxCMbQBUJKjTNbfuAlSeFdnt9F7mCQm5CRRIqrzNelwpE2bdmGjfw6mL25zWA0cb6BUplXN9ZRHIGraHq4RyMbUqBsibVhPhKy3rwA3qzCP9NbBhvVJORg7izBi6GojfjME0iJ8RdcfONttVLSObyoDuFpExxCK8AN4lkB5WCJnqZOQnu23Fy4GgZ7bGIvg59ukbKPkRM6EsZFf9PHulrisdw7hXsEGeJK9xcp4YMElFJ4OsTpFVKQSbnsloobcSsMCiZ5BgGKFcwlTHMnJOaEicL8uQft3BOcfOfAob44oFw3s1hcATPDAmUGrab0zeQgypm6RwcG67VhXvyNka4vV1jqbwfjwtGw0nQxgsGqb1jTFCOe4mdQbSxaeBLSj2mhCjTEcQoqgFeMLPHjlRYbss48Up"
	for i := 0; i < b.N; i++ {
		DefaultBoolean(str4096)
	}
}

// DefaultNumber

func TestDefaultNumberEmptyString(t *testing.T) {
	queryValue := ""
	defaultValue := 1
	expected := 1

	received := DefaultNumber(queryValue, defaultValue)

	if received != expected {
		t.Errorf("DefaultNumber(\"%v\", %d) failed. Expected %v, Received %v", queryValue, defaultValue, expected, received)
	} else {
		t.Logf("DefaultNumber(\"%v\", %d) passed.", queryValue, defaultValue)
	}
}

func TestDefaultNumberNotNumber(t *testing.T) {
	queryValue := "foo"
	defaultValue := 10
	expected := -1

	received := DefaultNumber(queryValue, defaultValue)

	if received != expected {
		t.Errorf("DefaultNumber(%v, %v) failed. Expected %v, Received %v", queryValue, defaultValue, expected, received)
	} else {
		t.Logf("DefaultNumber(%v, %d) passed.", queryValue, defaultValue)
	}
}

func TestDefaultNumberBig(t *testing.T) {
	queryValue := "123456789123456789"
	defaultValue := 10
	expected := -1

	received := DefaultNumber(queryValue, defaultValue)

	if received != expected {
		t.Errorf("DefaultNumber(%v, %v) failed. Expected %v, Received %v", queryValue, defaultValue, expected, received)
	} else {
		t.Logf("DefaultNumber(%v, %d) passed.", queryValue, defaultValue)
	}
}

func BenchmarkDefaultNumber8String(b *testing.B) {
	str8 := "12345678"
	for i := 0; i < b.N; i++ {
		DefaultNumber(str8, 10)
	}
}

func BenchmarkDefaultNumber1024String(b *testing.B) {
	str1024 := "RivB8d4iEuol9MLFllafqqAmdYqfXcUY5NzWhBrHabxP4RYceelq8F5NxYbSr56OyCOWgYfmquYsQSPyNyY5OXMjDHwmqDCocEKl8auiznqbZexviKvXDap62DuWemjQqpUpAntEzBahzulQmzYj01E1XCoUJvanvT4abnUktXr4XXUk4tzeOI78WdrD6FVUvN1eyMppfvRK1ScF0VKeCCIFLYEaG8b4BQYuU1HlcSuyE4fNSDpdwePwgYl3Fvb811uRhr4Fo374fArBTOc5YEyHMEO95Vo1dCjtItmPhFLdGcgrAENrPMZ5vh5J5llabH4dqE1N0Ej090w1POoGcUWeUsmYQ8enGSryLjcyuKLeAeLWi3SblLm4pBt5UnVaiHxNiZslfpLLPFKGD5j0DcrTCzX50DU6SumXXk8ON9K9KVy2q9GbrUvjvxi8V687mP37bax7UpV0j2dHkZ5XlG01GELPOpcEDoaltZ8XdzjXRjuRs1JuyaqoFIhp6bCP79bCOnCgBrNXxGzy0rgHo2USXD98NVlQW1v9H2zabVJK0UljMolXwAlBz7GRR4lQwuhge6OtujlJQK1Icg5mHpYyPaTk0Lj0luy0YwLbuwnbBjKGKneP3zLiwzuOyNpVDWwq8uuYhusQzS6arOFJxJTGbhOCNchTrDUbuwpvIuq0jXt8aIyXa1N9FiZeBegT3AVcpneMWNgC8fcvkaWG2Lo8DtqhCvp8hqSh6VcQDRa7XvG6pMcmRDXgkiyEKFt4Gatmiu7OcdNpEBOBTuFQ4EIieFImO5LIWeqiz9UWMQCMiHiKcOAoDXWasTPvh5j7aUhDbkr6McyeVUobxXixckzo1IY24SK2gBlEJvRbTwqImd4zAjgdbRaOL4vhTZjdmf4elvt5mq0gvI5pM8GULL3r4eKHLTJ1f9HpvUEePI4Dz8le4dgbWLnRwTLaxy0cxew0R18zapOLRWJCeSSuXUuQNRmIP2VXGGZrmtGBmTrGIf1UbvYxUw7GpamKt0W0"
	for i := 0; i < b.N; i++ {
		DefaultNumber(str1024, 10)
	}
}

func BenchmarkDefaultNumber4096String(b *testing.B) {
	str4096 := "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"
	for i := 0; i < b.N; i++ {
		DefaultNumber(str4096, 10)
	}
}

//// NumberOverMax

func TestNumberOverMaxNotNumber(t *testing.T) {
	value := "foo"
	expectedNumber := -1

	num, err := NumberOverMax(value)

	if num == expectedNumber && err != nil {
		t.Logf("NumberOverMax(\"foo\") passed.")
	} else {
		t.Errorf("NumberOverMax(\"foo\") failed. Expected %v, Received %v", expectedNumber, num)
	}
}

func TestNumberOverMaxLessThanZero(t *testing.T) {
	value := "-1"
	expectedNumber := -1
	expectedError := ErrorLessThanZero

	num, err := NumberOverMax(value)

	if num == expectedNumber && errors.Is(err, expectedError) {
		t.Logf("NumberOverMax(-1) passed.")
	} else {
		t.Errorf("NumberOverMax(-1) failed. Expected %v, Received %v", expectedNumber, num)
	}
}

func TestNumberOverMaxGreaterThanMax(t *testing.T) {
	value := "2147483647"
	expectedNumber := -1
	expectedError := ErrorGreaterThanMax

	num, err := NumberOverMax(value)

	if num == expectedNumber && errors.Is(err, expectedError) {
		t.Logf("NumberOverMax(\"2147483647\") passed.")
	} else {
		t.Errorf("NumberOverMax(\"2147483647\") failed. Expected %v, Received %v", expectedNumber, num)
	}
}

func TestNumberOverMax(t *testing.T) {
	value := "100"
	expectedNumber := 100
	var expectedError error

	num, err := NumberOverMax(value)

	if num == expectedNumber && err == expectedError {
		t.Logf("NumberOverMax(\"100\") passed.")
	} else {
		t.Errorf("NumberOverMax(\"100\") failed. Expected %v, Received %v", expectedNumber, num)
	}
}

// TODO: benchmark
func BenchmarkNumberOverMax(b *testing.B) {
	value := "100"
	for i := 0; i < b.N; i++ {
		NumberOverMax(value)
	}
}

func BenchmarkNumberOverMaxLessThanZero(b *testing.B) {
	value := "-1"
	for i := 0; i < b.N; i++ {
		NumberOverMax(value)
	}
}

func BenchmarkNumberOverMaxGreaterThanMax(b *testing.B) {
	value := "2147483647"
	for i := 0; i < b.N; i++ {
		NumberOverMax(value)
	}
}
