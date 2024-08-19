### Make public methods/functions easy to scan
Does your code have excessively long public methods or functions?
Encapsulate the steps in private methods/functions with the extract method (command+option+m) refactor.
The code below has some boring, distracting ceremony around creating a JSON string and turning it into an io.Reader so that we can POST it in an HTTP request.
```
func (ws *WidgetService) CreateWidget(name string) error {
	url := ws.baseURL + "/widgets"
	payload := []byte(`{"name": "` + name + `"}`)

	req, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer(payload),
	)
	//todo: handle codes, err etc
}
```
First, use the inline variable refactor (command+option+n) to put the payload into the buffer creation.
```
func (ws *WidgetService) CreateWidget(name string) error {
	url := ws.baseURL + "/widgets"
	req, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer([]byte(`{"name": "`+name+`"}`)),
	)
	// etc
}
```
Now, we can extract the creation of the JSON payload into a function using the extract method refactor (command+option+m) to remove the noise from the method.
```
func (ws *WidgetService) CreateWidget(name string) error {
	url := ws.baseURL + "/widgets"
	req, err := http.NewRequest(
		http.MethodPost,
		url,
		createWidgetPayload(name),
	)
	// etc
}
```
**Public methods and functions should describe what they do rather than how they do it**.
*Whenever I have to think to understand what the code is doing, I ask myself if I can refactor the code to make that understanding more immediately apparent*
-- Martin Fowler
This helps you understand the overall design better, and it then allows you to ask questions about responsibilities:
* Why does this method do X? Shouldn't that live in Y?
* Why does this method do so many tasks? Can we consolidate this elsewhere?
**Private functions and methods are great; they let you wrap up irrelevant how's into whats**.

Quite deliberately, as the writer of CreateWidget, I do not want the creation of a specific string to be an essential character in the narration of the method. It is distracting, irrelevant noise for the reader 99% of the time.
However, **if someone does care, you press command+b (or whatever "navigate to symbol" is for you) on createWidgetPayload ... and read it. Press command+left-arrow to go back again**.