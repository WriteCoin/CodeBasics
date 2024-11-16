package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func Greetings(name string) string {
	name = strings.Trim(name, " ")
	name = cases.Title(language.Russian).String(name)
	return fmt.Sprintf("Привет, %s!", name)
}

func ТестПеременных() {
	var (
		firstName string = "John"
		lastName  string = "Smith"
	)
	fmt.Println(firstName, lastName)
}

func TestIntToString() {
	fmt.Println(IntToString(5))
}

func ТестПриветствия() {
	fmt.Println(Greetings("Иван"))
}

func DomainForLocale(domain, locale string) string {
	if locale == "" {
		locale = "en"
	}
	// domain = strings.Trim(domain, ".")
	// domain = strings.Trim(domain, " ")
	return locale + "." + domain
}

// func DomainForLocale(domain, locale string) string {
//     if locale == "" {
//         locale = "en"
//     }
//     return fmt.Sprintf("%s.%s", locale, domain)
// }

func ТестDomainForLocale() {
	fmt.Println(DomainForLocale("site.com", ""))
	fmt.Println(DomainForLocale("site.com", "ru"))
}

func ModifySpaces(s, mode string) string {
	switch {
	case mode == "dash":
		return strings.ReplaceAll(s, " ", "-")
	case mode == "underscore":
		return strings.ReplaceAll(s, " ", "_")
	case mode == "unknown" || mode == "":
		return strings.ReplaceAll(s, " ", "*")
	default:
		return strings.ReplaceAll(s, " ", "*")
	}
}

// UserCreateRequest is a request to create a new user.
type UserCreateRequest struct {
	FirstName string
	Age       int
}

func Validate(req UserCreateRequest) string {
	if req.FirstName == "" || strings.Contains(req.FirstName, " ") || req.Age <= 0 || req.Age > 150 {
		return "invalid request"
	}
	return ""
}

const A = 20

func ErrorMessageToCode(msg string) int {
	switch msg {
	case "OK":
		return 0
	case "CANCELLED":
		return 1
	case "UNKNOWN":
		return 2
	default:
		return 2
	}
}

type Voicer interface {
	Voice() string
}

type Cat struct {
	// …
}

type Cow struct {
	// …
}

type Dog struct {
	// …
}

// BEGIN (write your solution here)
func (ca Cat) Voice() string {
	return "Мяу"
}

func (co Cow) Voice() string {
	return "Мууу"
}

func (do Dog) Voice() string {
	return "Гав"
}

func ТестМассивов(массив []int) []int {
	arr := [5]string{"Троянский конь"}

	fmt.Println(len(arr))
	fmt.Println(len(arr[0]))

	// массив := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	массив = append(массив, 11)
	// массив[0] = массив[len(массив) - 1]

	новыйМассив := make([]int, len(массив), len(массив))
	новыйМассив = массив

	return новыйМассив
}

func SafeWrite(nums [5]int, i, val int) [5]int {
	if i >= 0 && i < len(nums) {
		nums[i] = val
	}
	return nums
}

func Remove(nums []int, i int) []int {
	if i < 0 || i >= len(nums) {
		return nums
	}
	новыйИндекс := i + 1
	if новыйИндекс < len(nums) {
		nums[i] = nums[новыйИндекс]
	} else {
		return nums[:len(nums)-1]
	}
	return Remove(nums, новыйИндекс)
}

func RemoveInt64(nums []int64, i int) []int64 {
	if i < 0 || i >= len(nums) {
		return nums
	}

	nums[i] = nums[len(nums)-1]

	return nums[:len(nums)-1]
}

func Map(strs []string, mapFunc func(s string) string) []string {
	массив := make([]string, len(strs))

	for i, v := range strs {
		массив[i] = mapFunc(v)
	}

	return массив
}

func IntsCopy(src []int, maxLen int) []int {
	if maxLen <= 0 {
		return []int{}
	}
	if maxLen > len(src) {
		maxLen = len(src)
	}
	копия := make([]int, maxLen)
	copy(копия, src)
	return копия
}

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	for i := 0; i < len(userIDs); i++ {
		for j := i + 1; j < len(userIDs); j++ {
			if userIDs[i] == userIDs[j] {
				userIDs[j] = userIDs[len(userIDs)-1]

				userIDs = userIDs[:len(userIDs)-1]
			}
		}
	}

	sort.Slice(userIDs, func(i, j int) bool {
		return userIDs[i] < userIDs[j]
	})

	return userIDs
}

func ТестыМассивовИСрезов() {
	// ТестПриветствия()
	// ТестDomainForLocale()

	// nums := [5]int{}
	// fmt.Println(SafeWrite(nums, 30, 1))

	массив := []int{1, 2, 3, 4}
	массив = Remove(массив, 0)
	fmt.Println(массив)

	копияМассива := IntsCopy(массив, 1)
	fmt.Println(копияМассива)
	копияМассива[0] = 1
	fmt.Println(массив)

	массивНеуникальных := []int64{55, 2, 88, 33, 2, 2, 55, 103, 33, 88}
	массивУникальных := UniqueSortedUserIDs(массивНеуникальных)
	fmt.Println(массивУникальных)
}

func UniqueUserIDs(userIDs []int64) []int64 {
	отображение := map[int64]int{}
	сколькоПропустили := 0
	for i, v := range userIDs {
		_, существует := отображение[v]
		if !существует {
			отображение[v] = i - сколькоПропустили
		} else {
			сколькоПропустили++
		}
	}
	новыйМассив := make([]int64, len(отображение))
	for k, i := range отображение {
		новыйМассив[i] = k
	}
	return новыйМассив
}

func ТестUniqueUserIDs() {
	массив := []int64{55, 2, 88, 33, 2, 2, 55, 103, 33, 88}
	fmt.Println(UniqueUserIDs(массив))
}

func MostPopularWord(words []string) string {
	количестваСлов := make([]int, len(words))
	слова := make(map[string]int)
	for i, v := range words {
		_, существует := слова[v]
		if существует {
			количестваСлов[слова[v]]++
		} else {
			количестваСлов = append(количестваСлов, 0)
			слова[v] = i
		}
	}
	индексРезультата := 0
	частоеКоличество := 0
	for i, v := range количестваСлов {
		if v > частоеКоличество {
			частоеКоличество = v
			индексРезультата = i
		}
	}
	return words[индексРезультата]
}

func ТестMostPopularWord() {
	слова := []string{"a", "b", "c", "c", "d", "e", "e", "d"}
	результат := MostPopularWord(слова)
	fmt.Println(результат)
}

func nextASCII(b byte) byte {
	if b < 255 {
		return b + 1
	}
	return b
}

func prevASCII(b byte) byte {
	if b > 0 {
		return b - 1
	}
	return b
}

func shiftASCII(s string, step int) string {
	строка := make([]byte, len(s))
	if step > 256 {
		step = int(math.Abs(float64(step)))
	}
	for i := 0; i < len(s); i++ {
		строка[i] = s[i] + byte(step)
	}
	return string(строка)
}

func isASCII(s string) bool {
	for _, c := range s {
		if c > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func latinLetters(s string) string {
	var стройка strings.Builder

	for _, c := range s {
		if unicode.Is(unicode.Latin, c) {
			стройка.WriteString(string(c))
		}
	}

	return стройка.String()
}

func ТестСдвигаASCII() {
	строки := []string{"abc", "abc1", "bcd2", "hi", "abc", "abc"}
	сдвиги := []int{0, 1, -1, 10, 256, -511}

	for i := 0; i < len(строки); i++ {
		fmt.Println(строки[i], shiftASCII(строки[i], сдвиги[i]))
	}
}

func ТестASCII() {
	строки := []string{" abc1", "хай", "hello \U0001F970"}

	for i := 0; i < len(строки); i++ {
		fmt.Println(isASCII(строки[i]))
	}
}

func ТестСтрок() {
	s := "Привет"
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}

	ТестASCII()
}

func generateSelfStory(name string, age int, money float64) string {
	return fmt.Sprintf("Hello! My name is %s. I'm %d y.o. And I also have $%.2f in my wallet right now.", name, age, money)
}

func MergeNumberLists(numberLists ...[]int) []int {
	result := []int{}
	for _, срез := range numberLists {
		result = append(result, срез...)
	}
	return result
}

// Parent is a parent struct.
type Parent struct {
	Name     string
	Children []Child
}

// Child is a child struct.
type Child struct {
	Name string
	Age  int
}

func CopyParent(p *Parent) Parent {
	if p == nil {
		return Parent{}
	}
	дети := append([]Child{}, p.Children...)
	return Parent{
		Name:     p.Name,
		Children: дети,
	}
}

type Counter struct {
	Value int
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func (c *Counter) Inc(delta int) {
	if delta == 0 {
		delta = 1
	}
	c.Value += delta
}

func (c *Counter) Dec(delta int) {
	if delta == 0 {
		delta = 1
	}
	c.Value -= delta
	c.Value = Max(c.Value, 0)
}

func ТестСтруктур() {
	fmt.Println(MergeNumberLists([]int{1, 2}, []int{3}, []int{4}))

	cp := CopyParent(nil) // Parent{}

	fmt.Println(cp)

	p := &Parent{
		Name: "Harry",
		Children: []Child{
			{
				Name: "Andy",
				Age:  18,
			},
		},
	}
	копия := CopyParent(p)

	// при мутациях в копии "cp"
	// изначальная структура "p" не изменяется
	копия.Children[0] = Child{}
	fmt.Println(p.Children) // [{Andy 18}]
}

type кодОшибки string

type Person struct {
	Age uint8
}

type PersonList []Person

func (pl PersonList) GetAgePopularity() map[uint8]int {
	отображение := map[uint8]int{}
	for _, человек := range pl {
		отображение[человек.Age]++
	}
	return отображение
}

func ТестТипов() {
	ко := кодОшибки("внутренний")

	fmt.Println(reflect.TypeOf(ко))

	fmt.Println(reflect.TypeOf(string(ко)))
}

// CreateUserRequest is a request to create a new user.
type CreateUserRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

// validation errors
var (
	errEmailRequired                = errors.New("email is required")
	errPasswordRequired             = errors.New("password is required")
	errPasswordConfirmationRequired = errors.New("password confirmation is required")
	errPasswordDoesNotMatch         = errors.New("password does not match with the confirmation")
)

func DecodeAndValidateRequest(requestBody []byte) (CreateUserRequest, error) {
	запрос := CreateUserRequest{}

	ошибка := json.Unmarshal(requestBody, &запрос)

	if ошибка != nil {
		return запрос, ошибка
	}

	if strings.Trim(запрос.Email, " ") == "" {
		return CreateUserRequest{}, errEmailRequired
	}

	if strings.Trim(запрос.Password, " ") == "" {
		return CreateUserRequest{}, errPasswordRequired
	}

	if strings.Trim(запрос.PasswordConfirmation, " ") == "" {
		return CreateUserRequest{}, errPasswordConfirmationRequired
	}

	if запрос.Password != запрос.PasswordConfirmation {
		return CreateUserRequest{}, errPasswordDoesNotMatch
	}

	return запрос, ошибка
}

type nonCriticalError struct{}

func (e nonCriticalError) Error() string {
	return "validation error"
}

var (
	errBadConnection = errors.New("bad connection")
	errBadRequest    = errors.New("bad request")
)

const unknownErrorMsg = "unknown error"

func GetErrorMsg(err error) string {
	if errors.As(err, &nonCriticalError{}) {
		return ""
	}
	if errors.Is(err, errBadConnection) {
		return errBadConnection.Error()
	}
	if errors.Is(err, errBadRequest) {
		return errBadRequest.Error()
	}
	return unknownErrorMsg
	// if !errors.Is(err, errBadConnection) && !errors.Is(err, errBadRequest) {
	// 	return unknownErrorMsg
	// }
	// return string(err)
	// return fmt.Errorf("%w", err).Error()
	// return err.Error()

	// if err.Error() == errBadConnection.Error() || err.Error() == errBadRequest.Error() {
	// 	return fmt.Errorf("%w", err).Error()
	// }
	// return unknownErrorMsg
}

// MergeDictsJob is a job to merge dictionaries into a single dictionary.
type MergeDictsJob struct {
	Dicts      []map[string]string
	Merged     map[string]string
	IsFinished bool
}

// errors
var (
	errNotEnoughDicts = errors.New("at least 2 dictionaries are required")
	errNilDict        = errors.New("nil dictionary")
)

func ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error) {
	job.IsFinished = true
	if len(job.Dicts) < 2 {
		return job, errNotEnoughDicts
	}
	job.Merged = make(map[string]string)
	for _, d := range job.Dicts {
		if d == nil {
			return job, errNilDict
		}
		for k, v := range d {
			job.Merged[k] = v
		}
	}
	return job, nil
}

func ТестВалидацииИОшибок() {
	байты := []byte("{\"email\":\"test\",\"password\":\"test\",\"password_confirmation\":\"test\"}")
	fmt.Println(DecodeAndValidateRequest(байты))

	ошибки := []error{
		errors.New("bad connection"),
		errors.New("bad request"),
		nonCriticalError{},
		errors.New("random error"),
	}

	for _, ошибка := range ошибки {
		fmt.Println(GetErrorMsg(ошибка))
	}
}

func MaxSum(nums1, nums2 []int) []int {
	l1 := len(nums1)
	s1 := 0
	s2 := 0
	nums := append(nums1, nums2...)
	for i, v := range nums {
		if i < l1 {
			s1 += v
		} else {
			s2 += v
		}
	}
	if s1 > s2 {
		return nums1
	} else if s1 < s2 {
		return nums2
	} else {
		return nums1
	}
}

func сумма(nums []int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	return s
}

func ПоточныйMaxSum(nums1, nums2 []int) []int {
	s1 := 0
	s2 := 0

	go func() { s1 = сумма(nums1) }()
	go func() { s2 = сумма(nums2) }()

	time.Sleep(100 * time.Millisecond)

	if s1 > s2 {
		return nums1
	} else if s1 < s2 {
		return nums2
	} else {
		return nums1
	}
}

func SumWorker(numsCh chan []int, sumCh chan int) {
	for nums := range numsCh {
		s := 0
		for _, v := range nums {
			s += v
		}
		sumCh <- s
		// go func(nums []int) {
		// 	for _, v := range nums {
		// 		s += v
		// 	}
		// 	sumCh <- s
		// }(nums)
	}
}

func printer(msgCh chan string) {
	// читаем из канала, пока он открыт
	for msg := range msgCh {
		fmt.Println(msg)
	}

	fmt.Println("printer has finished")
}

func ТестПотоков() {
	// создаем канал, в который будем отправлять сообщения
	msgCh := make(chan string)

	// вызываем функцию асинхронно в горутине
	go printer(msgCh)

	msgCh <- "hello"
	msgCh <- "concurrent"
	msgCh <- "world"

	// закрываем канал
	close(msgCh)

	// и ждем, пока printer закончит работу
	time.Sleep(100 * time.Millisecond)

	numsCh := make(chan []int)
	sumCh := make(chan int)

	go SumWorker(numsCh, sumCh)

	numsCh <- []int{1, 2, 3, 4}
	// numsCh <- nil
	// numsCh <- []int{}

	результат := <-sumCh

	fmt.Println(результат)
}

func main() {
	// ТестUniqueUserIDs()
	// ТестMostPopularWord()

	// ТестСтрок()
	// ТестСтруктур()
	// ТестТипов()

	// ТестВалидацииИОшибок()

	ТестПотоков()
}
