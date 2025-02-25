# ToDo API
## Beschreibung
Dies ist eine sichere ToDo List – API, auf der ein Frontend aufgebaut werden kann. Die API stellt sicher, dass ToDos nur vom jeweiligen Nutzer eingesehen werden können, der sie erstellt hat. Nutzer können ihre ToDos mit anderen Nutzern teilen. Die API wurde dementsprechend erweitert.

## User Stories
### Login
- **Beschreibung**: Als Benutzer möchte ich mich einloggen, damit nur ich meine ToDos sehen kann.
### ToDo erstellen
- **Beschreibung**: Als Benutzer möchte ich ein neues ToDo mit Titel und Text erstellen und in der Datenbank speichern.
### ToDo auflisten
- **Beschreibung**: Als Benutzer möchte ich alle meine ToDos in einer Liste sehen.
### ToDo löschen
- **Beschreibung**: Als Benutzer möchte ich ein ToDo aus der Liste löschen, wenn ich es nicht mehr benötige.
### Erledigt markieren
- **Beschreibung**: Als Benutzer möchte ich mein ToDo als erledigt markieren.

## Backlog:
### ToDo aktualisieren
- **Beschreibung**: Als Benutzer möchte ich ein bestehendes ToDo ändern können.
### ToDo verschieben
- **Beschreibung**: Als Benutzer möchte ich meine ToDos neu anordnen.
### ToDo Kategorien
- **Beschreibung**: Als Benutzer möchte ich meine ToDos kategorisieren.
### ToDo teilen
- **Beschreibung**: Als Benutzer möchte ich mein ToDo mit einem anderen Benutzer teilen, damit dieser auch meine ToDos sehen und als erledigt markieren kann.

## Installation
1. Klone das Repository
```bash
    git clone https://github.com/iribuda/todo-api-go.git
```
2. `db.sql` soll in DBMS gelaufen werden, um die Datenbank zu erstellen
3. Starte die Anwendung, die auf (http://localhost:8080) gelaufen wird
```bash
  go run cmd/main.go
```

## API Endpunkte
[Postman Collection](https://www.postman.com/descent-module-geoscientist-24846769/workspace/public/collection/18343508-4db2a613-28e9-49db-8ce5-b2ecf41b4b17?action=share&creator=18343508)

### Registrierung
- **POST /register**
  - Nutzerdaten: `{ username, email, password }`
  - Antwort: `{ message: 'Registered successfully'  }`

### Anmelden
- **POST /login**
  - Nutzerdaten: `{ email, password }`
  - Antwort: `{ token }`

### ToDos
- **GET /tasks**
  - Beschreibung: Alle Aufgaben des eingeloggten Nutzers abrufen.
  - Header: `Authorization: <token>`
  - Antwort: `[{ id, title, text, deadline, categoryId, category, done }]`

- **POST /tasks**
  - Beschreibung: Eine neue Aufgabe erstellen.
  - Header: `Authorization: <token>`
  - Nutzerdaten: `{ title, text }`
  - Antwort: `{ id, title, text, deadline, categoryId, category, done }`

- **PUT /tasks/:id**
  - Beschreibung: Eine bestehende Aufgabe aktualisieren.
  - Header: `Authorization: <token>`
  - Nutzerdaten: `{ title, text, deadline, categoryId, category }`
  - Antwort: `{ id, title, text, deadline, categoryId, category, done }`

- **DELETE /tasks/:id**
  - Beschreibung: Eine Aufgabe löschen.
  - Header: `Authorization: <token>`
  - Antwort: `{ message: 'Task deleted successfully' }`

- **POST /tasks/:id/complete**
  - Beschreibung: Eine Aufgabe als erledigt markieren.
  - Header: `Authorization: <token>`
  - Antwort: `{ id, title, text, deadline, categoryId, category, done }`

- **POST /tasks/:id/share**
  - Beschreibung:Eine Aufgabe mit einem anderen Nutzer teilen.
  - Header: `Authorization: <token>`
  - Nutzerdaten: `{ sharedUserID   }`
  - Antwort: `{ message: 'Task shared successfully' }`
 

## Datenbank
![image](https://github.com/iribuda/todo-api-go/assets/87514795/95b0d74d-3de9-4062-8792-cf5b13a7960c)
