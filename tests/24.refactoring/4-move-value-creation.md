### Move value creation to construction time

Methods often have to create value and use them.

```
type WidgetService struct {
	baseURL string
	client  *http.Client
}

func NewWidgetService(baseURL string) *WidgetService {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	return &WidgetService{baseURL: baseURL, client: &client}
}

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

A refactoring technique you could apply here is, **if a value is being created that is not dependant on the arguments to the method, then you can instead create a field in your type and calculate it in your constructor function**.

```
type WidgetService struct {
	client          *http.Client
	createWidgetURL string
}

func NewWidgetService(baseURL string) *WidgetService {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	return &WidgetService{
		createWidgetURL: baseURL + "/widgets",
		client:          &client,
	}
}

func (ws *WidgetService) CreateWidget(name string) error {
	req, err := http.NewRequest(
		http.MethodPost,
		ws.createWidgetURL,
		createWidgetPayload(name),
	)
	// etc
}
```

By moving them to construction time, you can simplify your methods.

These continuous minor improvements are vital to the long-term health of a codebase.
