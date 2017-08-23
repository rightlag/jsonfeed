> Parse and validate JSON feeds

[![Build Status](https://img.shields.io/travis/rightlag/jsonfeed.svg?style=flat-square)](https://travis-ci.org/rightlag/jsonfeed) [![Coverage Status](https://img.shields.io/coveralls/rightlag/jsonfeed.svg?style=flat-square)](https://coveralls.io/github/rightlag/jsonfeed)

# Usage

The JSON Feed below has exactly 5 issues:

- Root document is missing required properties `version` and `title`
- `items[0]` is missing required property `id`
- `items[0].attachments[0]` is missing required properties `url` and `mime_type`

See the JSON Feed Version 1 [Specification](https://jsonfeed.org/version/1) for more information regarding validation.

```json
{
    "version": "",
    "title": "",
    "home_page_url": "http://flyingmeat.com/blog/",
    "feed_url": "http://flyingmeat.com/blog/feed.json",
    "description": "News from your friends at Flying Meat.",
    "author": {
        "name": "Gus Mueller"
    },
    "items": [
        {
            "id": "",
            "title": "Acorn and Sierra Compatibility",
            "content_html": "<p>macOS Sierra is just around the corner, and if you&#39;re running the beta or developer seeds of it you&#39;re of course going to want to know if your favorite application, <a href=\"http://flyingmeat.com/acorn/\">Acorn 5</a>, is compatible with it.</p>\n<p>And it is of course.</p>\n<strike>There is at least one Sierra issue we are aware of (in beta 3, and it&#39;s totally Apple&#39;s fault). Exporting deep images (aka, 16 bits per component) is currently broken in the developer and public betas. I&#39;ve filed a radar with Apple and this is such a serious oversight on their part, that I&#39;m sure it&#39;ll be fixed pretty soon (#27285115 ImageIO problemo).</strike>\n\n<p><strong>Update August 1st, 2016:</strong> Apple has fixed this bug in beta 4, and we are no longer aware of any issues with Acorn and 10.12 Sierra.</p>\n<p>If you&#39;re running the Sierra betas and you encounter any crashes, bugs, or other issues let us <a href=\"mailto:support@flyingmeat.com\">know right away</a>! We want to make sure we don&#39;t miss anything.</p>\n",
            "date_published": "2016-07-22T09:02:14-07:00",
            "url": "http://flyingmeat.com/blog/archives/2016/7/acorn_and_sierra_compatibility.html",
            "attachments": [
                {
                    "url": "",
                    "mime_type": ""
                }
            ]
        }
    ]
}
```

The CLI parses and validates any given JSON Feed document. If the JSON document is malformed, an error is thrown.

    $ jsonfeed -file /path/to/file

```
❯ jsonfeed -file ~/feed.json
2017/08/22 21:08:19 missing required property `version`
2017/08/22 21:08:19 missing required property `title`
2017/08/22 21:08:19 item 1 missing required property `id`
2017/08/22 21:08:19 item 1 attachment 1 missing required property `url`
2017/08/22 21:08:19 item 1 attachment 1 missing required property `mime_type`
```

## Testing

	$ go test ./... -v

## License

[Jason Walsh](https://twitter.com/rightlag) &copy; [MIT](LICENSE)
