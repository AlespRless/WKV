package main

import (
	"fmt"
	"wkv/storage"
)

func main() {
	s := storage.New()
	if v, has := s.GetIfPresent("阿巴阿巴"); !has {
		fmt.Printf("没这个值\n")
		s.Put("阿巴阿巴", "差不多得了")
	} else {
		fmt.Printf("阿巴阿巴的值为: %v\n", v)
	}
	if v, has := s.GetIfPresent("阿巴阿巴"); !has {
		fmt.Printf("没这个值\n")
		s.Put("阿巴阿巴", "差不多得了")
	} else {
		fmt.Printf("阿巴阿巴的值为: %v\n", v)
	}
	s.Delete("阿巴阿巴")
	if v, has := s.GetIfPresent("阿巴阿巴"); !has {
		fmt.Printf("没这个值\n")
		s.Put("阿巴阿巴", "差不多得了")
	} else {
		fmt.Printf("阿巴阿巴的值为: %v\n", v)
	}
}
