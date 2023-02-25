package service

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/hexiaopi/drawio-store/internal/entity"
)

const (
	DefaultImageData = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAHkAAAA9CAYAAACJM8YzAAAAAXNSR0IArs4c6QAAA1d0RVh0bXhmaWxlACUzQ214ZmlsZSUyMGhvc3QlM0QlMjJlbWJlZC5kaWFncmFtcy5uZXQlMjIlMjBtb2RpZmllZCUzRCUyMjIwMjItMDMtMTFUMTUlM0EwNyUzQTU4LjExOVolMjIlMjBhZ2VudCUzRCUyMjUuMCUyMChXaW5kb3dzJTIwTlQlMjAxMC4wJTNCJTIwV2luNjQlM0IlMjB4NjQpJTIwQXBwbGVXZWJLaXQlMkY1MzcuMzYlMjAoS0hUTUwlMkMlMjBsaWtlJTIwR2Vja28pJTIwQ2hyb21lJTJGOTkuMC40ODQ0LjUxJTIwU2FmYXJpJTJGNTM3LjM2JTIwRWRnJTJGOTkuMC4xMTUwLjM2JTIyJTIwZXRhZyUzRCUyMlpYTDNxTnFfRmk2bko1dWxPTnp4JTIyJTIwdmVyc2lvbiUzRCUyMjE3LjEuMiUyMiUyMHR5cGUlM0QlMjJlbWJlZCUyMiUzRSUzQ2RpYWdyYW0lMjBpZCUzRCUyMmVtWVJTRERSa0RQMllkM0VaeTNyJTIyJTIwbmFtZSUzRCUyMlBhZ2UtMSUyMiUzRWpaTEJjb1FnRElhZmhydksxTzFldTkzdVhucnkwRE9WVkpnaWNSQlg3ZE5YUzZneU81M3BpZkFsSkg4U0dEJTJCMTA4V0pUcjJpQk1PS1RFNk1QN09peU11c1hJNlZ6SUU4SGc0Qk5FN0xnTElOVlBvTDZHV2tnNWJRRXd2SUl4cXZ1eFRXYUMzVVBtSENPUnpUc0E4MGFkVk9OSkJFcktDcWhibW5iMXA2UlYwVTVjYXZvQnNWSyUyQmZsTVhqZVJmM1pPQndzMWJOb0lYaGFFZE9RaGw0SmllTU84VFBqSjRmb2c5Vk9KekRyV05PSnZmemhKY205bjJNVFVhc0Q2JTJGJTJCVklUWnhFMlpJY3V5U2prcDdxRHBSciUyRmR4MlQ3alQ4cTNacm5saTNsZmszVGR3SG1ZZG9nMFhBQmI4RzVlUXNoN2ZBZ3Y2T3ZrQmFrYXQwWGtjWVpxdDRTU21LRGRONyUyQlp0NTRYZzlxTzEyM2VQNzdkZiUyQmJuYnclM0QlM0QlM0MlMkZkaWFncmFtJTNFJTNDJTJGbXhmaWxlJTNFngzGMAAAAPVJREFUeF7t04EJwEAMw8Bk/6G/FEqHeJ03kIR3Zs7Y1Qb2jXyOzrdW3t0R+da6H5fIlwd+8UQWOWAggOjJIgcMBBA9WeSAgQCiJ4scMBBA9GSRAwYCiJ4scsBAANGTRQ4YCCB6ssgBAwFETxY5YCCA6MkiBwwEED1Z5ICBAKInixwwEED0ZJEDBgKInixywEAA0ZNFDhgIIHqyyAEDAURPFjlgIIDoySIHDAQQPVnkgIEAoieLHDAQQPRkkQMGAoieLHLAQADRk0UOGAggerLIAQMBRE8WOWAggOjJIgcMBBA9WeSAgQCiJ4scMBBA/J8cYE0jPh5C7gE8XAF0AAAAAElFTkSuQmCC"
)

func GetDrawImage(name string) (image []byte, err error) {
	file, err := os.ReadFile("data/" + name)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func CreateDrawImage(name string) error {
	draw := entity.DrawImage{
		Alt:  name,
		Data: DefaultImageData,
	}
	return UpdateDraw(draw)
}

func DeleteDrawImage(name string) error {
	_, err := os.Stat("data/" + name)
	if err != nil {
		return err
	}
	return os.Remove("data/" + name)
}

func ListDrawImages() (images []entity.DrawImage, err error) {
	files, err := os.ReadDir("data")
	if err != nil {
		return nil, err
	}
	for _, v := range files {
		info, err := v.Info()
		if err != nil {
			return nil, err
		}
		// file, err := fs.ReadFile(Images, "data/"+v.Name())
		// if err != nil {
		// 	return nil, err
		// }
		images = append(images, entity.DrawImage{
			Alt:       v.Name(),
			Data:      "http://localhost:8080/api/v1/image/" + v.Name(),
			CreatedAt: info.ModTime(),
			Size:      info.Size(),
		})
	}
	return
}

func UpdateDraw(draw entity.DrawImage) error {
	encoded := string(draw.Data[22:])
	fmt.Println(encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return err
	}
	name := "data/" + draw.Alt
	return os.WriteFile(name, decoded, os.ModePerm)
}
