package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jimmyfielding/maps-api-project/pkg/api/v1beta1"
	"github.com/spf13/cobra"
)

func NewCmdGenerate(out, errOut io.Writer) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "generate",
		Run: func(cmd *cobra.Command, args []string) {
			if err := RunGenerate(cmd); err != nil {
				fmt.Println(err)
			}
		},
	}

	cmd.PersistentFlags().StringP("file", "f", "", "--file metadata.csv")

	return cmd
}

func RunGenerate(cmd *cobra.Command) error {
	file, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}

	defer f.Close()
	metadata := []v1beta1.ImageMetadata{}
	var i v1beta1.ImageMetadata
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		str := strings.Replace(record[0], "T", " ", 1)
		str = strings.Replace(str, "Z", "", 1)
		t, err := time.Parse("2006-01-02 15:04:05", str)
		if err != nil {
			return err
		}

		lat, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return err
		}

		lng, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return err
		}

		i = v1beta1.ImageMetadata{
			Time:      &t,
			Latitude:  lat,
			Longitude: lng,
		}

		metadata = append(metadata, i)
	}

	titles, err := generateTitles(metadata)
	if err != nil {
		return err
	}

	if len(titles) == 0 {
		fmt.Println("nothing generated")
	}

	printTitles(titles)

	return nil
}

func generateTitles(metadata []v1beta1.ImageMetadata) ([]v1beta1.Title, error) {
	titles, _, err := client.GenerateTitles(metadata)
	if err != nil {
		return []v1beta1.Title{}, err
	}

	return titles, nil
}

func printTitles(titles []v1beta1.Title) {
	var b strings.Builder
	b.WriteString("Generated the following title suggestion: \n\n")
	for _, t := range titles {
		b.WriteString(fmt.Sprintf("\t%s\n", t))
	}
	b.WriteString("\n")
	fmt.Println(b.String())
}
