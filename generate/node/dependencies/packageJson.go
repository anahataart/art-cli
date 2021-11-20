package dependencies

import (
	"encoding/json"
	"log"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pterm/pterm"

	"golang.org/x/text/language"
)

var localizer *i18n.Localizer
var bundle *i18n.Bundle

func PackageJson(project string, main string, lang string) {
	// Configurações para internacionalização do software
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

	_, errorPath := os.Stat("./" + project)

	localizeConfigErrorPackageJson := i18n.LocalizeConfig{
		MessageID: "error_package_json",
	}
	localizeConfigSuccessPackageJson := i18n.LocalizeConfig{
		MessageID: "success_package_json",
	}
	resultErrorPackageJson, _ := localizer.Localize(&localizeConfigErrorPackageJson)
	resultSuccessPackageJson, _ := localizer.Localize(&localizeConfigSuccessPackageJson)

	if os.IsNotExist(errorPath) {
		pterm.Error.Println(resultErrorPackageJson)
	} else {

		f, err := os.Create("./" + project + "/package.json")

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		firstLine := "{"
		data := []byte(firstLine)

		_, err2 := f.Write(data)

		if err2 != nil {
			log.Fatal(err2)
		}

		line2 := "\n   \"name\":  \"" + project + "\",\n"
		data2 := []byte(line2)

		var idx int64 = int64(len(data))

		_, err3 := f.WriteAt(data2, idx)

		if err3 != nil {
			log.Fatal(err3)
		}

		pterm.Success.Println(resultSuccessPackageJson)
	}
}
