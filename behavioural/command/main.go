package main

import "fmt"

type somebuf struct {
	text string
}

type command interface {
	execute()
	snapshot() *somebuf
}

type setCommand struct {
	command
	buf     *somebuf
	save    somebuf
	newText string
}

func (cmd setCommand) execute() {
	cmd.buf.text = cmd.newText
	cmd.save = *cmd.buf
	fmt.Println("edit=" + cmd.snapshot().text)
}

func (cmd setCommand) snapshot() *somebuf {
	return &cmd.save
}

///// revert

type revertCommand struct {
	command
	buf *somebuf
}

func (cmd revertCommand) execute() {
	cmd.buf.text = cmd.snapshot().text
	fmt.Println("undo=" + cmd.snapshot().text)
}

///// delete

type deleteCommand struct {
	command
	buf  *somebuf
	save somebuf
}

func (cmd deleteCommand) execute() {
	cmd.buf.text = ""
	cmd.save = *cmd.buf
	fmt.Println("delete=" + cmd.snapshot().text)
}

func (cmd deleteCommand) snapshot() *somebuf {
	return &cmd.save
}

func main() {
	var mybuf somebuf

	var cmds []command

	cmds = append(cmds, setCommand{newText: "first", buf: &mybuf})
	cmds = append(cmds, setCommand{newText: "second", buf: &mybuf})
	cmds = append(cmds, deleteCommand{buf: &mybuf})
	cmds = append(cmds, setCommand{newText: "third", buf: &mybuf})

	for _, x := range cmds {
		x.execute()
	}

	revert := cmds[2]
	revert.execute()
	revert = cmds[0]
	revert.execute()
}
