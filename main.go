package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TwiN/go-color"
	"github.com/google/uuid"
)

type Task struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Done bool `json:"done"`
}
type List struct {
	Tasks []Task `json:"tasks"`
}
func main() {
	// commands:
	// kyoto add example
	// kyoto delete 07bb3347-e738-40a0-a78a-5e6de7df4cd7
	// kyoto delete all
	// kyoto list
	// kyoto done 07bb3347-e738-40a0-a78a-5e6de7df4cd7
	// kyoto done all
	// kyoto undone 07bb3347-e738-40a0-a78a-5e6de7df4cd7
	// kyoto undone all
	// kyoto help 
	tasks_path := "/home/user/tasks.json"
	if _, err := os.Stat(tasks_path); err != nil {
		task_list := List{
			Tasks: []Task{},
		}
		prettier, _ := json.MarshalIndent(task_list, "", "\t")
		f, err := os.OpenFile(tasks_path, os.O_CREATE|os.O_WRONLY, 0640)
		if err != nil {
			panic(err)
		}
		if _, err := f.WriteString(string(prettier)); err != nil {
			panic(err)
		}
		if err := f.Close(); err != nil {
			panic(err)
		}
	}
	if len(os.Args) > 1 {
		if os.Args[1] == "add" {
			if len(os.Args) == 3 {
				f, err := os.ReadFile(tasks_path)
				if err != nil {
				panic(err)
				}
				var list_task List
				err = json.Unmarshal(f, &list_task)
				if err != nil {
					panic(err)
				}
				id := uuid.New()
				list_task.Tasks = append(list_task.Tasks, Task{ID: id.String(), Name: os.Args[2], Done: false})
				pretty, _ := json.MarshalIndent(list_task, "", "\t")
				f1, err := os.OpenFile(tasks_path, os.O_CREATE|os.O_WRONLY, 0640)
				if err != nil {
					panic(err)
				}
				if _, err := f1.WriteString(string(pretty)); err != nil {
					panic(err)
				}
				if err := f1.Close(); err != nil {
					panic(err)
				}
				fmt.Println("Added.")
			} else {
				fmt.Println("Too many arguments or missing something.")
			} 
		} else if os.Args[1] == "delete" {
			if len(os.Args) == 3 {
				if os.Args[2] == "all" {
					if err := os.Truncate(tasks_path, 0); err != nil {
						panic(err)
					}
					task_list := List{
						Tasks: []Task{},
					}
					prettier, _ := json.MarshalIndent(task_list, "", "\t")
					file_1, err := os.OpenFile(tasks_path, os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						panic(err)
					}
					if _, err := file_1.WriteString(string(prettier)); err != nil {
						panic(err)
					}
					if err := file_1.Close(); err != nil {
						panic(err)
					}
					fmt.Println("All tasks got deleted.")
				} else {
					f, err := os.ReadFile(tasks_path)
					if err != nil {
					panic(err)
					}
					var list_task List
					err = json.Unmarshal(f, &list_task)
					if err != nil {
						panic(err)
					}
					index := 0
					for c, i := range list_task.Tasks {
						if i.ID == os.Args[2] {
							index = c
						}
					}
					if index != 0 {
						list_task.Tasks = append(list_task.Tasks[:index], list_task.Tasks[index+1:]...)
					}
					if err := os.Truncate(tasks_path, 0); err != nil {
						panic(err)
					}
					prettier, _ := json.MarshalIndent(list_task, "", "\t")
					file_2, err := os.OpenFile(tasks_path, os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						panic(err)
					}
					if _, err := file_2.WriteString(string(prettier)); err != nil {
						panic(err)
					}
					if err := file_2.Close(); err != nil {
						panic(err)
					}
					fmt.Printf("Task %s got deleted.", os.Args[2])
				}
			} else {
				fmt.Println("Too many arguments or missing something.")
			}
		} else if os.Args[1] == "list" {
			f, err := os.ReadFile(tasks_path)
			if err != nil {
				panic(err)
			}
			var list_task List
			err = json.Unmarshal(f, &list_task)
			if err != nil {
				panic(err)
			}
			if len(list_task.Tasks) == 0 {
				fmt.Println("No tasks.")
			}
			for c, i := range list_task.Tasks {
				jumpline := ""
				if len(list_task.Tasks) > 1 {
					jumpline = "\n\n"
				}
				if len(list_task.Tasks) == c+1 {
					jumpline = ""
				}
				normal := fmt.Sprintf("%s%s%s", color.Red, fmt.Sprintf("%v", i.Done), color.Reset)
				if i.Done {
					normal = fmt.Sprintf("%s%s%s", color.Green, fmt.Sprintf("%v", i.Done), color.Reset)
				}
				fmt.Printf("ID: %s\nName: %s\nDone: %s%s", color.Gray + i.ID + color.Reset, i.Name, normal, jumpline)
			}
		} else if os.Args[1] == "done" {
			if len(os.Args) == 3 {
				if os.Args[2] == "all" {
					f, err := os.ReadFile(tasks_path)
					if err != nil {
						panic(err)
					}
					var list_task List
					err = json.Unmarshal(f, &list_task)
					if err != nil {
						panic(err)
					}
					for c := range list_task.Tasks {
						list_task.Tasks[c].Done = true
					}
					if err := os.Truncate(tasks_path, 0); err != nil {
						panic(err)
					}
					prettier, _ := json.MarshalIndent(list_task, "", "\t")
					file_2, err := os.OpenFile(tasks_path, os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						panic(err)
					}
					if _, err := file_2.WriteString(string(prettier)); err != nil {
						panic(err)
					}
					if err := file_2.Close(); err != nil {
						panic(err)
					}
					fmt.Println("All tasks are done.")
				} else {
					f, err := os.ReadFile(tasks_path)
					if err != nil {
						panic(err)
					}
					var list_task List
					err = json.Unmarshal(f, &list_task)
					if err != nil {
						panic(err)
					}
					for c := range list_task.Tasks {
						if os.Args[2] == list_task.Tasks[c].ID {
							list_task.Tasks[c].Done = true
						}
					}
					if err := os.Truncate(tasks_path, 0); err != nil {
						panic(err)
					}
					prettier, _ := json.MarshalIndent(list_task, "", "\t")
					file_2, err := os.OpenFile(tasks_path, os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						panic(err)
					}
					if _, err := file_2.WriteString(string(prettier)); err != nil {
						panic(err)
					}
					if err := file_2.Close(); err != nil {
						panic(err)
					}
					fmt.Printf("Task %s is done.\n", os.Args[2])
				}
			} else {
				fmt.Println("Too many arguments or missing something.")
			}
		} else if os.Args[1] == "help" {
			fmt.Printf("commands:\nkyoto add example\nkyoto delete 07bb3347-e738-40a0-a78a-5e6de7df4cd7\nkyoto delete all\nkyoto list\nkyoto done 07bb3347-e738-40a0-a78a-5e6de7df4cd7\nkyoto done all\nkyoto undone 07bb3347-e738-40a0-a78a-5e6de7df4cd7\nkyoto undone all\nkyoto help")
		} else if os.Args[1] == "undone" {
			if len(os.Args) == 3 {
				if os.Args[2] == "all" {
					f, err := os.ReadFile(tasks_path)
					if err != nil {
						panic(err)
					}
					var list_task List
					err = json.Unmarshal(f, &list_task)
					if err != nil {
						panic(err)
					}
					for c := range list_task.Tasks {
						list_task.Tasks[c].Done = false
					}
					if err := os.Truncate(tasks_path, 0); err != nil {
						panic(err)
					}
					prettier, _ := json.MarshalIndent(list_task, "", "\t")
					file_2, err := os.OpenFile(tasks_path, os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						panic(err)
					}
					if _, err := file_2.WriteString(string(prettier)); err != nil {
						panic(err)
					}
					if err := file_2.Close(); err != nil {
						panic(err)
					}
					fmt.Println("All tasks got undoned.")
				} else {
					f, err := os.ReadFile(tasks_path)
					if err != nil {
						panic(err)
					}
					var list_task List
					err = json.Unmarshal(f, &list_task)
					if err != nil {
						panic(err)
					}
					for c := range list_task.Tasks {
						if os.Args[2] == list_task.Tasks[c].ID {
							list_task.Tasks[c].Done = false
						}
					}
					if err := os.Truncate(tasks_path, 0); err != nil {
						panic(err)
					}
					prettier, _ := json.MarshalIndent(list_task, "", "\t")
					file_2, err := os.OpenFile(tasks_path, os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						panic(err)
					}
					if _, err := file_2.WriteString(string(prettier)); err != nil {
						panic(err)
					}
					if err := file_2.Close(); err != nil {
						panic(err)
					}
					fmt.Printf("Task %s got undoned.\n", os.Args[2])
				}
			} else {
				fmt.Println("Too many arguments or missing something.")
			}
		} else {
			fmt.Println("Nothing else matters.")
		}
	} else {
		fmt.Println("Missing everything.")
	}
}