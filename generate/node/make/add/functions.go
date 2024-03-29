package add

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/project-barca/barca-cli/utils/file"
	"github.com/pterm/pterm"

	"golang.org/x/text/language"
)

func init() {
	log.SetPrefix("FUNCTIONS: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

var localizer *i18n.Localizer
var bundle *i18n.Bundle

func FunctionJS(lang, directory, collection, database, method, visibility, types string, hidden bool, inputs []string) {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("translate/en.json")
	bundle.LoadMessageFile("translate/fr.json")
	bundle.LoadMessageFile("translate/pt-BR.json")

	switch lang {
	case "portugues-brasileiro":
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	case "english":
		localizer = i18n.NewLocalizer(bundle, language.English.String(), language.BrazilianPortuguese.String(), language.French.String())
	case "francais":
		localizer = i18n.NewLocalizer(bundle, language.French.String(), language.English.String(), language.BrazilianPortuguese.String())
	default:
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	}

	_, errorPath := os.Stat("./" + directory)

	if hidden != true {

		switch method {
		case "create":
			localizeConfigErrorModelFunctionsJavacript := i18n.LocalizeConfig{
				MessageID: "error_functions_javascript",
			}
			localizeConfigSuccessFunctionsJavascript := i18n.LocalizeConfig{
				MessageID: "success_functions_javascript",
			}
			resultErrorModelFunctionsJavascript, _ := localizer.Localize(&localizeConfigErrorModelFunctionsJavacript)
			resultSuccessFunctionsJavascript, _ := localizer.Localize(&localizeConfigSuccessFunctionsJavascript)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsJavascript)
			} else {
				path := directory + "/js"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".js") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".js", "function insert"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfunction insert" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "  "+input+",")
							}
						}
						functiondata = append(functiondata, "};")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".js", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".js", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função CREATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".js")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "function insert" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("  " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".js", "};", file.GetNumberLines(path+"/"+collection+".js")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".js", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsJavascript)
			}

		case "read":
			localizeConfigErrorModelFunctionsJavacript := i18n.LocalizeConfig{
				MessageID: "error_functions_javascript",
			}
			localizeConfigSuccessFunctionsJavascript := i18n.LocalizeConfig{
				MessageID: "success_functions_javascript",
			}
			resultErrorModelFunctionsJavascript, _ := localizer.Localize(&localizeConfigErrorModelFunctionsJavacript)
			resultSuccessFunctionsJavascript, _ := localizer.Localize(&localizeConfigSuccessFunctionsJavascript)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsJavascript)
			} else {
				path := directory + "/js"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".js") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".js", "function read"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfunction read" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "  "+input+",")
							}
						}
						functiondata = append(functiondata, "};")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".js", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".js", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função READ")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".js")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "function read" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("  " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".js", "};", file.GetNumberLines(path+"/"+collection+".js")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".js", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}
				pterm.Success.Println(resultSuccessFunctionsJavascript)
			}

		case "delete":
			localizeConfigErrorModelFunctionsJavacript := i18n.LocalizeConfig{
				MessageID: "error_functions_javascript",
			}
			localizeConfigSuccessFunctionsJavascript := i18n.LocalizeConfig{
				MessageID: "success_functions_javascript",
			}
			resultErrorModelFunctionsJavascript, _ := localizer.Localize(&localizeConfigErrorModelFunctionsJavacript)
			resultSuccessFunctionsJavascript, _ := localizer.Localize(&localizeConfigSuccessFunctionsJavascript)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsJavascript)
			} else {
				path := directory + "/js"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".js") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".js", "function delete"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfunction delete" + collection + "(id) {",
							"};",
						}
						file, err := os.OpenFile(path+"/"+collection+".js", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(file)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						file.Close()
					} else {
						fmt.Println("Já existe está função DELETE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".js")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "function delete" + collection + "(id) {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "id, \n};"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)
					if err3 != nil {
						log.Fatal(err3)
					}
				}
				pterm.Success.Println(resultSuccessFunctionsJavascript)
			}
		case "update":
			localizeConfigErrorModelFunctionsJavacript := i18n.LocalizeConfig{
				MessageID: "error_functions_javascript",
			}
			localizeConfigSuccessFunctionsJavascript := i18n.LocalizeConfig{
				MessageID: "success_functions_javascript",
			}
			resultErrorModelFunctionsJavascript, _ := localizer.Localize(&localizeConfigErrorModelFunctionsJavacript)
			resultSuccessFunctionsJavascript, _ := localizer.Localize(&localizeConfigSuccessFunctionsJavascript)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsJavascript)
			} else {
				path := directory + "/js"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".js") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".js", "function update"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfunction update" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "  "+input+",")
							}
						}
						functiondata = append(functiondata, "};")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".js", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".js", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função UPDATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".js")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "function update" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("  " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".js", "};", file.GetNumberLines(path+"/"+collection+".js")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".js", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsJavascript)
			}
		default:
			// barcaSQL
		}
	}

}

func FunctionPHP(lang, directory, collection, database, method, visibility, types string, hidden bool, inputs []string) {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("translate/en.json")
	bundle.LoadMessageFile("translate/fr.json")
	bundle.LoadMessageFile("translate/pt-BR.json")

	switch lang {
	case "portugues-brasileiro":
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	case "english":
		localizer = i18n.NewLocalizer(bundle, language.English.String(), language.BrazilianPortuguese.String(), language.French.String())
	case "francais":
		localizer = i18n.NewLocalizer(bundle, language.French.String(), language.English.String(), language.BrazilianPortuguese.String())
	default:
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	}

	_, errorPath := os.Stat("./" + directory)

	if hidden != true {

		switch method {
		case "create":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/php"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".php") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".php", "function insert"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"<?php\n\nfunction insert" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += "  $" + input + ", "
								functiondata = append(functiondata, "  $"+input+", ")
							}
						}
						functiondata = append(functiondata, "};")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".php", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".php", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função CREATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".php")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "<?php\n\nfunction insert" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += "$" + input + ", "
							_, err3 := f.WriteString("  $" + input + ";\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".php", "};", file.GetNumberLines(path+"/"+collection+".php")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".php", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsPHP)
			}

		case "read":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/php"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".php") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".php", "function read"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"<?php\n\nfunction read" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += "$" + input + ", "
								functiondata = append(functiondata, "  $"+input+",")
							}
						}
						functiondata = append(functiondata, "};")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".php", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".php", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função READ")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".php")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "<?php\n\nfunction read" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += "$" + input + ", "
							_, err3 := f.WriteString("  $" + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".php", "};", file.GetNumberLines(path+"/"+collection+".php")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".php", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}
				pterm.Success.Println(resultSuccessFunctionsPHP)
			}

		case "delete":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/php"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".php") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".php", "function delete"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"<?php\n\nfunction delete" + collection + "($id) {",
							"};",
						}
						file, err := os.OpenFile(path+"/"+collection+".php", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(file)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						file.Close()
					} else {
						fmt.Println("Já existe está função DELETE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".php")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "<?php\n\nfunction delete" + collection + "($id) {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "$id; \n};"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)
					if err3 != nil {
						log.Fatal(err3)
					}
				}
				pterm.Success.Println(resultSuccessFunctionsPHP)
			}
		case "update":
			localizeConfigErrorModelFunctionsJavacript := i18n.LocalizeConfig{
				MessageID: "error_functions_javascript",
			}
			localizeConfigSuccessFunctionsJavascript := i18n.LocalizeConfig{
				MessageID: "success_functions_javascript",
			}
			resultErrorModelFunctionsJavascript, _ := localizer.Localize(&localizeConfigErrorModelFunctionsJavacript)
			resultSuccessFunctionsJavascript, _ := localizer.Localize(&localizeConfigSuccessFunctionsJavascript)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsJavascript)
			} else {
				path := directory + "/php"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".php") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".php", "function update"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfunction update" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += "$" + input + ", "
								functiondata = append(functiondata, "  $"+input+",")
							}
						}
						functiondata = append(functiondata, "};")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".php", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".php", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função UPDATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".php")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "function update" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("  $" + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".php", "};", file.GetNumberLines(path+"/"+collection+".php")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".php", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsJavascript)
			}
		default:
			// barcaSQL
		}
	}

}

func FunctionRuby(lang, directory, collection, database, method, visibility, types string, hidden bool, inputs []string) {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("translate/en.json")
	bundle.LoadMessageFile("translate/fr.json")
	bundle.LoadMessageFile("translate/pt-BR.json")

	switch lang {
	case "portugues-brasileiro":
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	case "english":
		localizer = i18n.NewLocalizer(bundle, language.English.String(), language.BrazilianPortuguese.String(), language.French.String())
	case "francais":
		localizer = i18n.NewLocalizer(bundle, language.French.String(), language.English.String(), language.BrazilianPortuguese.String())
	default:
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	}

	_, errorPath := os.Stat("./" + directory)

	if hidden != true {

		switch method {
		case "create":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/ruby"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".rb") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".rb", "def insert"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\ndef insert" + collection + "\n"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "  "+input+", ")
							}
						}
						functiondata = append(functiondata, "end")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".rb", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".rb", "insert"+collection, "insert"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função CREATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".rb")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "def insert" + collection + "\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("  " + input + ";\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".rb", "end", file.GetNumberLines(path+"/"+collection+".rb")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".rb", "insert"+collection, "insert"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsPHP)
			}

		case "read":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/rb"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".rb") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".rb", "def read"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\ndef read" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "  "+input+",")
							}
						}
						functiondata = append(functiondata, "end")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".rb", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".rb", "read"+collection, "read"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função READ")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".rb")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "def read" + collection + "\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("  " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".rb", "end", file.GetNumberLines(path+"/"+collection+".rb")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".rb", "read"+collection, "read"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}
				pterm.Success.Println(resultSuccessFunctionsPHP)
			}

		case "delete":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/ruby"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".rb") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".rb", "def delete"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\ndef delete" + collection + "(id)",
							"end",
						}
						file, err := os.OpenFile(path+"/"+collection+".rb", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(file)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						file.Close()
					} else {
						fmt.Println("Já existe está função DELETE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".rb")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "def delete" + collection + "(id)\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "id; \nend"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)
					if err3 != nil {
						log.Fatal(err3)
					}
				}
				pterm.Success.Println(resultSuccessFunctionsPHP)
			}
		case "update":
			localizeConfigErrorModelFunctionsJavacript := i18n.LocalizeConfig{
				MessageID: "error_functions_javascript",
			}
			localizeConfigSuccessFunctionsJavascript := i18n.LocalizeConfig{
				MessageID: "success_functions_javascript",
			}
			resultErrorModelFunctionsJavascript, _ := localizer.Localize(&localizeConfigErrorModelFunctionsJavacript)
			resultSuccessFunctionsJavascript, _ := localizer.Localize(&localizeConfigSuccessFunctionsJavascript)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsJavascript)
			} else {
				path := directory + "/ruby"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".rb") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".rb", "def update"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\ndef update" + collection}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "  "+input+",")
							}
						}
						functiondata = append(functiondata, "end")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".rb", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".rb", "update"+collection, "update"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função UPDATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".rb")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "def update" + collection + "\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("  " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".rb", "end", file.GetNumberLines(path+"/"+collection+".rb")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".rb", "update"+collection, "update"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsJavascript)
			}
		default:
			// barcaSQL
		}
	}

}

func FunctionPython(lang, directory, collection, database, method, visibility, types string, hidden bool, inputs []string) {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("translate/en.json")
	bundle.LoadMessageFile("translate/fr.json")
	bundle.LoadMessageFile("translate/pt-BR.json")

	switch lang {
	case "portugues-brasileiro":
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	case "english":
		localizer = i18n.NewLocalizer(bundle, language.English.String(), language.BrazilianPortuguese.String(), language.French.String())
	case "francais":
		localizer = i18n.NewLocalizer(bundle, language.French.String(), language.English.String(), language.BrazilianPortuguese.String())
	default:
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	}

	_, errorPath := os.Stat("./" + directory)

	if hidden != true {

		switch method {
		case "create":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/python"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".py") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".py", "def insert"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\ndef insert" + collection + "():"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += "    " + input + ", "
								functiondata = append(functiondata, "    "+input+", ")
							}
						}

						fileCurrent, err := os.OpenFile(path+"/"+collection+".py", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".py", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função CREATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".py")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "def insert" + collection + "():\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("    " + input + ";\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".py", "", file.GetNumberLines(path+"/"+collection+".py")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".py", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsPHP)
			}

		case "read":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/python"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".py") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".py", "def read"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\ndef read" + collection + "():"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "    "+input+",")
							}
						}

						fileCurrent, err := os.OpenFile(path+"/"+collection+".py", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".py", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função READ")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".py")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "def read" + collection + "():\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("    " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".py", "", file.GetNumberLines(path+"/"+collection+".py")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".py", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}
				pterm.Success.Println(resultSuccessFunctionsPHP)
			}

		case "delete":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/python"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".py") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".py", "def delete"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\ndef delete" + collection + "(id):",
							"    print(id)",
						}
						file, err := os.OpenFile(path+"/"+collection+".py", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(file)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						file.Close()
					} else {
						fmt.Println("Já existe está função DELETE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".py")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "def delete" + collection + "(id):\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "    print(id)"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)
					if err3 != nil {
						log.Fatal(err3)
					}
				}
				pterm.Success.Println(resultSuccessFunctionsPHP)
			}
		case "update":
			localizeConfigErrorModelFunctionsJavacript := i18n.LocalizeConfig{
				MessageID: "error_functions_javascript",
			}
			localizeConfigSuccessFunctionsJavascript := i18n.LocalizeConfig{
				MessageID: "success_functions_javascript",
			}
			resultErrorModelFunctionsJavascript, _ := localizer.Localize(&localizeConfigErrorModelFunctionsJavacript)
			resultSuccessFunctionsJavascript, _ := localizer.Localize(&localizeConfigSuccessFunctionsJavascript)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsJavascript)
			} else {
				path := directory + "/python"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".py") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".py", "def update"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\ndef update" + collection + "():"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "    "+input+",")
							}
						}

						fileCurrent, err := os.OpenFile(path+"/"+collection+".py", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".py", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função UPDATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".py")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "def update" + collection + "():\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("    " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".py", "", file.GetNumberLines(path+"/"+collection+".py")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".py", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsJavascript)
			}
		default:
			// barcaSQL
		}
	}

}

func FunctionRust(lang, directory, collection, database, method, visibility, types string, hidden bool, inputs []string) {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("translate/en.json")
	bundle.LoadMessageFile("translate/fr.json")
	bundle.LoadMessageFile("translate/pt-BR.json")

	switch lang {
	case "portugues-brasileiro":
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	case "english":
		localizer = i18n.NewLocalizer(bundle, language.English.String(), language.BrazilianPortuguese.String(), language.French.String())
	case "francais":
		localizer = i18n.NewLocalizer(bundle, language.French.String(), language.English.String(), language.BrazilianPortuguese.String())
	default:
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	}

	_, errorPath := os.Stat("./" + directory)

	if hidden != true {

		switch method {
		case "create":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/rs"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".rs") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".rs", "fn insert"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfn insert" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "    "+input+", ")
							}
						}
						functiondata = append(functiondata, "}")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".rs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".rs", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função CREATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".rs")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "\nfn insert" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("    " + input + ";\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".rs", "}", file.GetNumberLines(path+"/"+collection+".rs")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".rs", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsPHP)
			}

		case "read":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/rs"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".rs") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".rs", "fn read"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfn read" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "    "+input+",")
							}
						}
						functiondata = append(functiondata, "}")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".rs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".rs", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função READ")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".rs")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "\nfn read" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("    " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".rs", "}", file.GetNumberLines(path+"/"+collection+".rs")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".rs", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}
				pterm.Success.Println(resultSuccessFunctionsPHP)
			}

		case "delete":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/rs"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".rs") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".rs", "fn delete"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfn delete" + collection + "(id) {",
							"}",
						}
						file, err := os.OpenFile(path+"/"+collection+".rs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(file)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						file.Close()
					} else {
						fmt.Println("Já existe está função DELETE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".rs")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "\nfn delete" + collection + "(id) {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "id; \n}"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)
					if err3 != nil {
						log.Fatal(err3)
					}
				}
				pterm.Success.Println(resultSuccessFunctionsPHP)
			}
		case "update":
			localizeConfigErrorModelFunctionsJavacript := i18n.LocalizeConfig{
				MessageID: "error_functions_javascript",
			}
			localizeConfigSuccessFunctionsJavascript := i18n.LocalizeConfig{
				MessageID: "success_functions_javascript",
			}
			resultErrorModelFunctionsJavascript, _ := localizer.Localize(&localizeConfigErrorModelFunctionsJavacript)
			resultSuccessFunctionsJavascript, _ := localizer.Localize(&localizeConfigSuccessFunctionsJavascript)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsJavascript)
			} else {
				path := directory + "/rs"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".rs") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".rs", "function update"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfn update" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "    "+input+",")
							}
						}
						functiondata = append(functiondata, "}")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".rs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".rs", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função UPDATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".rs")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "fn update" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("    " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".rs", "}", file.GetNumberLines(path+"/"+collection+".rs")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".rs", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsJavascript)
			}
		default:
			// barcaSQL
		}
	}

}

func FunctionCplus(lang, directory, collection, database, method, visibility, types string, hidden bool, inputs []string) {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("translate/en.json")
	bundle.LoadMessageFile("translate/fr.json")
	bundle.LoadMessageFile("translate/pt-BR.json")

	switch lang {
	case "portugues-brasileiro":
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	case "english":
		localizer = i18n.NewLocalizer(bundle, language.English.String(), language.BrazilianPortuguese.String(), language.French.String())
	case "francais":
		localizer = i18n.NewLocalizer(bundle, language.French.String(), language.English.String(), language.BrazilianPortuguese.String())
	default:
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	}

	_, errorPath := os.Stat("./" + directory)

	if hidden != true {

		switch method {
		case "create":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/c"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".c") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".c", "void insert"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nvoid insert" + collection + "() \n{"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "    "+input+", ")
							}
						}
						functiondata = append(functiondata, "}")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".c", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".c", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função CREATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".c")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "\nvoid insert" + collection + "()\n{\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("    " + input + ";\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".c", "}", file.GetNumberLines(path+"/"+collection+".c")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".c", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsPHP)
			}

		case "read":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/c"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".c") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".c", "void read"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nvoid read" + collection + "()\n{"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "    "+input+",")
							}
						}
						functiondata = append(functiondata, "}")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".c", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".c", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função READ")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".c")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "\nvoid read" + collection + "()\n{\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("    " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".c", "}", file.GetNumberLines(path+"/"+collection+".c")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".c", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}
				pterm.Success.Println(resultSuccessFunctionsPHP)
			}

		case "delete":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/c"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".c") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".c", "void delete"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nvoid delete" + collection + "(int id)\n{",
							"}",
						}
						file, err := os.OpenFile(path+"/"+collection+".c", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(file)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						file.Close()
					} else {
						fmt.Println("Já existe está função DELETE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".c")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "\nvoid delete" + collection + "(int id)\n{\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "id; \n}"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)
					if err3 != nil {
						log.Fatal(err3)
					}
				}
				pterm.Success.Println(resultSuccessFunctionsPHP)
			}
		case "update":
			localizeConfigErrorModelFunctionsJavacript := i18n.LocalizeConfig{
				MessageID: "error_functions_javascript",
			}
			localizeConfigSuccessFunctionsJavascript := i18n.LocalizeConfig{
				MessageID: "success_functions_javascript",
			}
			resultErrorModelFunctionsJavascript, _ := localizer.Localize(&localizeConfigErrorModelFunctionsJavacript)
			resultSuccessFunctionsJavascript, _ := localizer.Localize(&localizeConfigSuccessFunctionsJavascript)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsJavascript)
			} else {
				path := directory + "/c"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".c") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".c", "function update"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nvoid update" + collection + "()\n{"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "    "+input+",")
							}
						}
						functiondata = append(functiondata, "}")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".c", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".c", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função UPDATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".c")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "void update" + collection + "()\n{\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("    " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".c", "}", file.GetNumberLines(path+"/"+collection+".c")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".c", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsJavascript)
			}
		default:
			// barcaSQL
		}
	}

}

func FunctionCsharp(lang, directory, collection, database, method, visibility, types string, hidden bool, inputs []string) {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("translate/en.json")
	bundle.LoadMessageFile("translate/fr.json")
	bundle.LoadMessageFile("translate/pt-BR.json")

	switch lang {
	case "portugues-brasileiro":
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	case "english":
		localizer = i18n.NewLocalizer(bundle, language.English.String(), language.BrazilianPortuguese.String(), language.French.String())
	case "francais":
		localizer = i18n.NewLocalizer(bundle, language.French.String(), language.English.String(), language.BrazilianPortuguese.String())
	default:
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	}

	_, errorPath := os.Stat("./" + directory)

	if hidden != true {

		switch method {
		case "create":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/c#"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".cs") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".cs", "public void insert"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\npublic void insert" + collection + "() \n{"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "    "+input+", ")
							}
						}
						functiondata = append(functiondata, "}")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".cs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".cs", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função CREATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".cs")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "\npublic void insert" + collection + "()\n{\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("    " + input + ";\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".cs", "}", file.GetNumberLines(path+"/"+collection+".cs")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".cs", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsPHP)
			}

		case "read":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/cs"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".cs") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".cs", "public void read"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\npublic void read" + collection + "()\n{"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "    "+input+",")
							}
						}
						functiondata = append(functiondata, "}")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".cs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".cs", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função READ")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".cs")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "\npublic void read" + collection + "()\n{\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("    " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".cs", "}", file.GetNumberLines(path+"/"+collection+".cs")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".cs", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}
				pterm.Success.Println(resultSuccessFunctionsPHP)
			}

		case "delete":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/cs"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".cs") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".cs", "public void delete"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\npublic void delete" + collection + "(int id)\n{",
							"}",
						}
						file, err := os.OpenFile(path+"/"+collection+".cs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(file)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						file.Close()
					} else {
						fmt.Println("Já existe está função DELETE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".cs")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "\npublic void delete" + collection + "(int id)\n{\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "id; \n}"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)
					if err3 != nil {
						log.Fatal(err3)
					}
				}
				pterm.Success.Println(resultSuccessFunctionsPHP)
			}
		case "update":
			localizeConfigErrorModelFunctionsJavacript := i18n.LocalizeConfig{
				MessageID: "error_functions_javascript",
			}
			localizeConfigSuccessFunctionsJavascript := i18n.LocalizeConfig{
				MessageID: "success_functions_javascript",
			}
			resultErrorModelFunctionsJavascript, _ := localizer.Localize(&localizeConfigErrorModelFunctionsJavacript)
			resultSuccessFunctionsJavascript, _ := localizer.Localize(&localizeConfigSuccessFunctionsJavascript)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsJavascript)
			} else {
				path := directory + "/cs"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".cs") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".cs", "public void update"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\npublic void update" + collection + "()\n{"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "    "+input+",")
							}
						}
						functiondata = append(functiondata, "}")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".cs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".cs", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função UPDATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".cs")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "public void update" + collection + "()\n{\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("    " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".cs", "}", file.GetNumberLines(path+"/"+collection+".cs")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".cs", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsJavascript)
			}
		default:
			// barcaSQL
		}
	}

}

func FunctionKotlin(lang, directory, collection, database, method, visibility, types string, hidden bool, inputs []string) {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("translate/en.json")
	bundle.LoadMessageFile("translate/fr.json")
	bundle.LoadMessageFile("translate/pt-BR.json")

	switch lang {
	case "portugues-brasileiro":
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	case "english":
		localizer = i18n.NewLocalizer(bundle, language.English.String(), language.BrazilianPortuguese.String(), language.French.String())
	case "francais":
		localizer = i18n.NewLocalizer(bundle, language.French.String(), language.English.String(), language.BrazilianPortuguese.String())
	default:
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	}

	_, errorPath := os.Stat("./" + directory)

	if hidden != true {

		switch method {
		case "create":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/kt"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".kt") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".kt", "fun insert"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfun insert" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "    "+input+", ")
							}
						}
						functiondata = append(functiondata, "}")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".kt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".kt", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função CREATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".kt")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "\nfun insert" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("    " + input + ";\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".kt", "}", file.GetNumberLines(path+"/"+collection+".kt")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".kt", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsPHP)
			}

		case "read":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/kt"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".kt") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".kt", "fun read"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfun read" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "    "+input+",")
							}
						}
						functiondata = append(functiondata, "}")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".kt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".kt", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função READ")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".kt")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "\nfun read" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("    " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".kt", "}", file.GetNumberLines(path+"/"+collection+".kt")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".kt", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}
				pterm.Success.Println(resultSuccessFunctionsPHP)
			}

		case "delete":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/kt"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".kt") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".kt", "fun delete"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfun delete" + collection + "(id: Int) {",
							"}",
						}
						file, err := os.OpenFile(path+"/"+collection+".kt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(file)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						file.Close()
					} else {
						fmt.Println("Já existe está função DELETE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".kt")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "\nfun delete" + collection + "(id: Int) {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "id; \n}"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)
					if err3 != nil {
						log.Fatal(err3)
					}
				}
				pterm.Success.Println(resultSuccessFunctionsPHP)
			}
		case "update":
			localizeConfigErrorModelFunctionsJavacript := i18n.LocalizeConfig{
				MessageID: "error_functions_javascript",
			}
			localizeConfigSuccessFunctionsJavascript := i18n.LocalizeConfig{
				MessageID: "success_functions_javascript",
			}
			resultErrorModelFunctionsJavascript, _ := localizer.Localize(&localizeConfigErrorModelFunctionsJavacript)
			resultSuccessFunctionsJavascript, _ := localizer.Localize(&localizeConfigSuccessFunctionsJavascript)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsJavascript)
			} else {
				path := directory + "/kt"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".kt") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".kt", "fun update"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfun update" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "    "+input+",")
							}
						}
						functiondata = append(functiondata, "}")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".kt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".kt", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função UPDATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".kt")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "fun update" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("    " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".kt", "}", file.GetNumberLines(path+"/"+collection+".kt")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".kt", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsJavascript)
			}
		default:
			// barcaSQL
		}
	}

}

func FunctionSwift(lang, directory, collection, database, method, visibility, types string, hidden bool, inputs []string) {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("translate/en.json")
	bundle.LoadMessageFile("translate/fr.json")
	bundle.LoadMessageFile("translate/pt-BR.json")

	switch lang {
	case "portugues-brasileiro":
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	case "english":
		localizer = i18n.NewLocalizer(bundle, language.English.String(), language.BrazilianPortuguese.String(), language.French.String())
	case "francais":
		localizer = i18n.NewLocalizer(bundle, language.French.String(), language.English.String(), language.BrazilianPortuguese.String())
	default:
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	}

	_, errorPath := os.Stat("./" + directory)

	if hidden != true {

		switch method {
		case "create":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/swift"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".swift") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".swift", "func insert"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfunc insert" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "  "+input+", ")
							}
						}
						functiondata = append(functiondata, "}")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".swift", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".swift", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função CREATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".swift")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "\nfunc insert" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += "_ " + input + ", "
							_, err3 := f.WriteString("  " + input + ";\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".swift", "}", file.GetNumberLines(path+"/"+collection+".swift")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".swift", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsPHP)
			}

		case "read":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/swift"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".swift") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".swift", "func read"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfunc read" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += "_ " + input + ", "
								functiondata = append(functiondata, "  "+input+",")
							}
						}
						functiondata = append(functiondata, "}")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".swift", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".swift", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função READ")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".swift")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "\nfunc read" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += "_ " + input + ", "
							_, err3 := f.WriteString("  " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".swift", "}", file.GetNumberLines(path+"/"+collection+".swift")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".swift", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}
				pterm.Success.Println(resultSuccessFunctionsPHP)
			}

		case "delete":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/swift"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".swift") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".swift", "func delete"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfunc delete" + collection + "(id: Int) {",
							"}",
						}
						file, err := os.OpenFile(path+"/"+collection+".swift", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(file)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						file.Close()
					} else {
						fmt.Println("Já existe está função DELETE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".swift")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "\nfunc delete" + collection + "(id: Int) {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "id; \n}"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)
					if err3 != nil {
						log.Fatal(err3)
					}
				}
				pterm.Success.Println(resultSuccessFunctionsPHP)
			}
		case "update":
			localizeConfigErrorModelFunctionsJavacript := i18n.LocalizeConfig{
				MessageID: "error_functions_javascript",
			}
			localizeConfigSuccessFunctionsJavascript := i18n.LocalizeConfig{
				MessageID: "success_functions_javascript",
			}
			resultErrorModelFunctionsJavascript, _ := localizer.Localize(&localizeConfigErrorModelFunctionsJavacript)
			resultSuccessFunctionsJavascript, _ := localizer.Localize(&localizeConfigSuccessFunctionsJavascript)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsJavascript)
			} else {
				path := directory + "/swift"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".swift") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".swift", "func update"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfunc update" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += "_ " + input + ", "
								functiondata = append(functiondata, "  "+input+",")
							}
						}
						functiondata = append(functiondata, "}")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".swift", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".swift", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função UPDATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".swift")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "func update" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += "_ " + input + ", "
							_, err3 := f.WriteString("  " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".swift", "}", file.GetNumberLines(path+"/"+collection+".swift")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".swift", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsJavascript)
			}
		default:
			// barcaSQL
		}
	}

}

func FunctionScala(lang, directory, collection, database, method, visibility, types string, hidden bool, inputs []string) {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("translate/en.json")
	bundle.LoadMessageFile("translate/fr.json")
	bundle.LoadMessageFile("translate/pt-BR.json")

	switch lang {
	case "portugues-brasileiro":
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	case "english":
		localizer = i18n.NewLocalizer(bundle, language.English.String(), language.BrazilianPortuguese.String(), language.French.String())
	case "francais":
		localizer = i18n.NewLocalizer(bundle, language.French.String(), language.English.String(), language.BrazilianPortuguese.String())
	default:
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	}

	_, errorPath := os.Stat("./" + directory)

	if hidden != true {

		switch method {
		case "create":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/scala"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".sc") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".sc", "def insert"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\ndef insert" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "    "+input+", ")
							}
						}
						functiondata = append(functiondata, "}")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".sc", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".sc", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função CREATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".sc")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "\ndef insert" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("    " + input + ";\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".sc", "}", file.GetNumberLines(path+"/"+collection+".sc")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".sc", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsPHP)
			}

		case "read":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/sc"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".sc") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".sc", "def read"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\ndef read" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "    "+input+",")
							}
						}
						functiondata = append(functiondata, "}")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".sc", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".sc", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função READ")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".sc")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "\ndef read" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("    " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".sc", "}", file.GetNumberLines(path+"/"+collection+".sc")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".sc", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}
				pterm.Success.Println(resultSuccessFunctionsPHP)
			}

		case "delete":
			localizeConfigErrorModelFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "error_functions_php",
			}
			localizeConfigSuccessFunctionsPHP := i18n.LocalizeConfig{
				MessageID: "success_functions_php",
			}
			resultErrorModelFunctionsPHP, _ := localizer.Localize(&localizeConfigErrorModelFunctionsPHP)
			resultSuccessFunctionsPHP, _ := localizer.Localize(&localizeConfigSuccessFunctionsPHP)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsPHP)
			} else {
				path := directory + "/sc"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".sc") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".sc", "def delete"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\ndef delete" + collection + "(id: Int) : Int = {",
							"}",
						}
						file, err := os.OpenFile(path+"/"+collection+".sc", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(file)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						file.Close()
					} else {
						fmt.Println("Já existe está função DELETE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".sc")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "\ndef delete" + collection + "(id: Int) {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "id; \n}"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)
					if err3 != nil {
						log.Fatal(err3)
					}
				}
				pterm.Success.Println(resultSuccessFunctionsPHP)
			}
		case "update":
			localizeConfigErrorModelFunctionsJavacript := i18n.LocalizeConfig{
				MessageID: "error_functions_javascript",
			}
			localizeConfigSuccessFunctionsJavascript := i18n.LocalizeConfig{
				MessageID: "success_functions_javascript",
			}
			resultErrorModelFunctionsJavascript, _ := localizer.Localize(&localizeConfigErrorModelFunctionsJavacript)
			resultSuccessFunctionsJavascript, _ := localizer.Localize(&localizeConfigSuccessFunctionsJavascript)

			if os.IsNotExist(errorPath) {
				pterm.Error.Println(resultErrorModelFunctionsJavascript)
			} else {
				path := directory + "/sc"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
				}

				if file.CheckIfExists(path+"/"+collection+".sc") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".sc", "func update"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfunc update" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "    "+input+",")
							}
						}
						functiondata = append(functiondata, "}")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".sc", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".sc", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função UPDATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".sc")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "def update" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("    " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".sc", "}", file.GetNumberLines(path+"/"+collection+".sc")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".sc", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsJavascript)
			}
		default:
			// barcaSQL
		}
	}

}
