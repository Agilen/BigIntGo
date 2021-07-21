package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	bigintegers "example.com/BigIntegerGo/BigIntegers"
)

func main() {

	{
		a := bigintegers.ReadHex("40D4ED6B22B4A26625AFFF98B70342C0742C4EE21087230415DF1B9348B28C94")
		b := bigintegers.ReadHex("1A98996C6EFBC1BC3C230BE9272861A04689D8D76C4F361DCD35972D469197B4")
		start := time.Now()
		for i := 0; i < 1000; i++ {

			if bigintegers.DelLeadZero(bigintegers.ToHex(bigintegers.LongAdd(a, b))) != "5B6D86D791B0642261D30B81DE2BA460BAB627B97CD65921E314B2C08F442448" {
				fmt.Println("Error")
				break
			}

		}
		duration := time.Since(start)
		fmt.Println("Average Add:", duration/1000)
	}
	{
		a := bigintegers.ReadHex("3AF01A7357B25888DD937053E63DF5BC8562ED86D24295AC8C491BF41E428869")
		b := bigintegers.ReadHex("FC31AC9F4BB19608207B449B6F318CE53ECEEEA214C2981971036F45F587932")
		start := time.Now()
		for i := 0; i < 1000; i++ {
			if bigintegers.DelLeadZero(bigintegers.ToHex(bigintegers.LongSub(a, b))) != "2B2CFFA962F73F285B8BBC0A2F4ADCEE3175FE9CB0F66C2AF538E4FFBEEA0F37" {
				fmt.Println("Error")
				break
			}
		}
		duration := time.Since(start)
		fmt.Println("Average Sub:", duration.Nanoseconds()/1000)
	}
	{
		a := bigintegers.ReadHex("1FA5C57629704BEC9142567B9ECFD3FADF4029E8171C39AB6A7F4BB5551D7AC1")
		b := bigintegers.ReadHex("75F4A39FEAE14D872A5E8374B27CB3FADE464D35AAE3D9B285478AE0563EEE6A")
		start := time.Now()
		for i := 0; i < 1000; i++ {
			if bigintegers.DelLeadZero(bigintegers.ToHex(bigintegers.LongMul(a, b))) != "E95017987491347D9BC2B9020D688F1E2EB93636C68200070F3C80FCD14BED1C22D14814FF44D6B28D7EFC15BD8BA1D9DF399AC8DFCA9F16018E577371241EA" {
				fmt.Println("Error")
				break
			}
		}
		duration := time.Since(start)
		fmt.Println("Average Mul:", duration.Nanoseconds()/1000)
	}
	{
		a := bigintegers.ReadHex("6FC747E8A92E7ADD219DA48AF56A378B7D484FF9E2CEC81C24970D982CD381EE3CEC65072296645350319B24752497AF4B06B81284F25927C3DC71EED5345CE7")
		b := bigintegers.ReadHex("181F440F6C8BF3FBBA82755EDA369685FEF7226AD6BDC38D3646E61D86084768")
		start := time.Now()
		for i := 0; i < 1000; i++ {
			q, r := bigintegers.LongDivMod(a, b)
			if bigintegers.DelLeadZero(bigintegers.ToHex(q)) != "4A24442CA54CD28B3B33B52431401635A6EAE51068339E9FDD89354916534C498" {
				fmt.Println("Error")
				break
			}
			if bigintegers.DelLeadZero(bigintegers.ToHex(r)) != "51265554E04C5E6E7C300D3FB117FBAB8D6F6E0206BED147ADE2606607E5727" {
				fmt.Println("Error")
				break
			}

		}
		duration := time.Since(start)
		fmt.Println("Average DivMod:", duration.Nanoseconds()/1000)
	}
	{
		a := bigintegers.ReadHex("FC31AC9F4BB19608207B449B6F318CE53ECEEEA214C2981971036F45F587932")
		b := bigintegers.ReadDec("100")
		n := bigintegers.ReadHex("3AF01A7357B25888DD937053E63DF5BC8562ED86D24295AC8C491BF41E428869")
		start := time.Now()
		for i := 0; i < 1000; i++ {
			if bigintegers.DelLeadZero(bigintegers.ToHex(bigintegers.LongModPowerBarrett(a, b, n))) != strings.ToUpper("B5526DFA1D89386721A537C8B2BA93052733FB23424BB8CBB78B19EB7D15071") {
				fmt.Println(bigintegers.DelLeadZero(bigintegers.ToHex(bigintegers.LongModPowerBarrett(a, b, n))))
				fmt.Println("Error")
				break
			}
		}
		duration := time.Since(start)
		fmt.Println("Average Pow:", duration.Nanoseconds()/1000)
	}

	// TestAdd()
	// TestSub()
	// TestMul()
	// TestCmp()
	// TestDiv()

}
func TestAdd() {
	fmt.Println("====Test Add====")
	q := ReadFile("test_add.arith")
	l := 0
	for i := 0; i < len(q); i += 4 {
		a := bigintegers.ReadHex(string(q[i]))
		b := bigintegers.ReadHex(string(q[i+1]))
		c := bigintegers.LongAdd(a, b)
		if bigintegers.DelLeadZero(bigintegers.ToHex(c)) == string(q[i+2]) {

			l++
		} else {
			fmt.Println(false)
			fmt.Println(bigintegers.DelLeadZero(bigintegers.ToHex(c)))
			fmt.Println(string(q[i+2]))
		}
	}
	if l == 20 {
		fmt.Println("Test passed")
	}
}
func TestSub() {
	fmt.Println("====Test Sub====")
	q := ReadFile("test_sub.arith")
	l := 0
	for i := 0; i < len(q); i += 4 {
		a := bigintegers.ReadHex(string(q[i]))
		b := bigintegers.ReadHex(string(q[i+1]))
		c := bigintegers.LongSub(a, b)
		if bigintegers.DelLeadZero(bigintegers.ToHex(c)) == string(q[i+2]) {

			l++
		} else {
			fmt.Println(false)
			fmt.Println(bigintegers.DelLeadZero(bigintegers.ToHex(c)))
			fmt.Println(string(q[i+2]))
		}
	}
	if l == 21 {
		fmt.Println("Test passed")
	}
}
func TestMul() {
	fmt.Println("====Test Mul====")
	q := ReadFile("test_mul.arith")
	l := 0
	for i := 0; i < len(q); i += 4 {
		a := bigintegers.ReadHex(string(q[i]))
		b := bigintegers.ReadHex(string(q[i+1]))
		c := bigintegers.LongMul(a, b)
		if bigintegers.DelLeadZero(bigintegers.ToHex(c)) == string(q[i+2]) {

			l++
		} else {
			fmt.Println(false)
			fmt.Println(bigintegers.DelLeadZero(bigintegers.ToHex(c)))
			fmt.Println(string(q[i+2]))
		}
	}
	if l == 20 {
		fmt.Println("Test passed")
	}
}
func TestCmp() {
	fmt.Println("====Test Cmp====")
	q := ReadFile("test_cmp.arith")
	l := 0

	for i := 0; i < len(q); i += 4 {
		a := bigintegers.ReadHex(string(q[i]))
		b := bigintegers.ReadHex(string(q[i+1]))
		c := bigintegers.LongCmp(a, b)
		n, _ := strconv.Atoi(q[i+2])
		if c == n {

			l++
		} else {
			fmt.Println(false)
			fmt.Println(c)
			fmt.Println(string(q[i+2]))
		}
	}
	if l == 20 {
		fmt.Println("Test passed")
	}
}
func TestDiv() {
	fmt.Println("====Test Div====")
	q := ReadFile("test_div.arith")
	l := 0

	for i := 0; i < len(q); i += 4 {
		a := bigintegers.ReadHex(string(q[i]))
		b := bigintegers.ReadHex(string(q[i+1]))
		c, cmod := bigintegers.LongDivMod(a, b)
		div := strings.Fields(q[i+2])
		if bigintegers.DelLeadZero(bigintegers.ToHex(c)) == string(div[0]) && bigintegers.DelLeadZero(bigintegers.ToHex(cmod)) == string(div[1]) {

			l++
		} else {
			fmt.Println(false)
			fmt.Println(
				(bigintegers.DelLeadZero(bigintegers.ToHex(c))),
				(bigintegers.DelLeadZero(bigintegers.ToHex(cmod))),
			)
			fmt.Println(string(q[i+2]))
		}
	}
	if l == 20 {
		fmt.Println("Test passed")
	}
}

func ReadFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var q []string
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		q = append(q, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return q
}
