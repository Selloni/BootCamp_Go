для тестинга на утечки будем использовать pprof
https://golang-blog.blogspot.com/2019/12/memory-leaks-investigation-pprof.html

код для создания файла для нахождения утечки
file, err := os.Create("profile_data.txt")
if err != nil {
log.Fatal(err)
}
defer file.Close()
pprof.Lookup("heap").WriteTo(file, 0)