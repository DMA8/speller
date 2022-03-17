Сервис получает сообщение из nats, отдает его спеллеру, а спеллер отдает исправленный вариант в storage.

Каждые n минут storage сохраняется в Dump/spellcheck.csv

К Storage можно обратиться по http

	GET .../read?name=word - отдает word и все связанные с ним опечатки
	
	POST .../add добавляет к существующему списку опечток указанные опечатки (в теле запроса должен быть json вида {"spellName":"csvName", "misSpells":["misspell1", "misspell2]})
	
	POST ../create создает связку слово-опечатка (в теле запроса должен быть json вида {"spellName":"csvName", "misSpells":["misspell1", "misspell2]})
	
	DELETE ../fullDelete?name=csvName полностью удаляет связку слово-опечатки
	
	DELETE ../delete удаляет указанные опечатки из списка. (в теле запроса должен быть json вида {"spellName":"csvName", "misSpells":["misspell1", "misspell2]})
	
Чтобы работал спеллер, необходимо в корень репозитория добавить директорию "datasets/ru" в которой поместить словарь частот "freq-dict.txt.gz" и "sentences.txt.gz". Также добавить папку "models" в которую поместить modelName.gz, что является предобученной моделью.
