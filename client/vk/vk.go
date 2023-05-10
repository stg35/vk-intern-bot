package vk

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/stg35/vk-intern-bot/errors"
)

type Client struct {
	host     string
	basePath string
	token    string
	group_id int
	client   http.Client
}

func New(host string, token string, group_id int) *Client {
	return &Client{
		host:     host,
		basePath: "method",
		token:    token,
		group_id: group_id,
		client:   http.Client{},
	}
}

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	defer errors.Wrap(errors.RequestFailed, err)
	query.Add("v", "5.131")
	query.Add("access_token", c.token)

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) getSessionData() (*GetLongPollServerResponse, error) {
	q := url.Values{}
	q.Add("group_id", strconv.Itoa(c.group_id))

	data, err := c.doRequest("groups.getLongPollServer", q)
	if err != nil {
		return nil, err
	}
	defer func() { err = errors.Wrap(errors.SessionDataFailed, err) }()

	var res GetLongPollServerResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetUpdates(offset int) ([]Update, error) {
	data, err := c.doLongPollRequest(offset)
	if err != nil {
		return nil, err
	}
	defer func() { err = errors.Wrap(errors.GetUpdatesFailed, err) }()

	var res GetUpdatesResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Updates, nil
}

func (c *Client) doLongPollRequest(offset int) ([]byte, error) {
	session_data, err := c.getSessionData()
	if err != nil {
		return nil, err
	}
	defer errors.Wrap(errors.LongPollRequestFailed, err)

	query := url.Values{}
	query.Add("act", "a_check")
	query.Add("key", session_data.Data.Key)
	query.Add("ts", strconv.Itoa(offset))
	query.Add("wait", "25")

	req, err := http.NewRequest(http.MethodGet, session_data.Data.Server, nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
