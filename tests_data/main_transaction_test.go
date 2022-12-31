package tests_data

import (
	"fiff_golang_draft/helper"
	"fiff_golang_draft/services"
	test1_buy "fiff_golang_draft/tests_data/test1-buy"
	test2_buy "fiff_golang_draft/tests_data/test2-buy"
	test3_buy "fiff_golang_draft/tests_data/test3-buy"
	test4_buy "fiff_golang_draft/tests_data/test4-buy"
	test5_buy "fiff_golang_draft/tests_data/test5-buy"
	test6_sell "fiff_golang_draft/tests_data/test6-sell"
	test7_sell "fiff_golang_draft/tests_data/test7-sell"

	"testing"
)

func TestMainTransaction(t *testing.T) {
	helper.TestingMode = true
	services.InitParamDb()
	test1_buy.Execute(t)
	test2_buy.Execute(t)
	test3_buy.Execute(t)
	test4_buy.Execute(t)
	test5_buy.Execute(t)
	test6_sell.Execute(t)
	test7_sell.Execute(t)

}
