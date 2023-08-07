**Git i18n demo**

Use gettext as i18n framework for go project.

    $ make
    $ ./go-i18n po/build/locale

## Mark your code

Import "i18n" package in your code, e.g.:

    import "github.com/jiangxin/go-i18n/i18n"

Setup locale dir and use default language, and returns a
locale instance.

	l := i18n.Setup("/opt/demo/share/locale", "demo")

Mark string using the wrappers provided by i18n package:

    fmt.Print(l.L_("Hello, world.\n"))


## Extract marked string into a template file: po/xx.pot

Extract marked strings in source code using `xgettext`.
But you can use the wrapper in Makefile:

    make pot

It will extract i18n strings into a template file: `po/go-i18n.pot`.

See file `makefile.i18n`. It is borrowed from Git project.


## Prepare to translate

As a l10n translater, prepare your language file for translate.
You can use `msgmerge`, but you can also use the wrapper in Makefile:

    make po-init PO_FILE=po/zh_CN.po

This command will create a language file for Simplified Chinese translation.

You can also create language file for other languages, such as:

    make po-init PO_FILE=po/zh_TW.po
    make po-init PO_FILE=po/fr.po

Please note: use locale name as your langue file. See your locale code by
running:

    locale -a | sed 's/\..*$//g' | sort


## Start to translate your language file

It is recommended that you use Emacs with po mode as the editor for po files.


## Compile language files

You can use `msgfmt` to compile `.po` files into compiled `.mo` files,
but you can also use wapper in Makefile:

    make po/zh_CN.mo


## Install

Install compiled language files in proper location with your program,
and your program can speak multiple languages.
