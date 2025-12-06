package main

import "strings"

func check(f_name string, l_name string, tickit_no uint, date string, email string, reminiingseaats uint) (bool, bool, bool) {
	isvalidname := len(f_name) > 2 && len(l_name) > 2
	isvalidateemail := strings.Contains(email, "@")
	isvalidateremaining := tickit_no < reminiingseaats

	return isvalidateemail, isvalidateremaining, isvalidname
}
