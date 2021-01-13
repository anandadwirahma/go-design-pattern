package main

/*
	Prototype interface
*/

type inode interface {
	print(string)
	clone() inode
}
