<a name="readme-top"></a>

<div align="center">

[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]

# Go Insert Locale

Insert locale data into your `locales` directory.

[Report an issue](https://github.com/AlejandroSuero/go-insert-locale/issues/new?assignees=&labels=bug&projects=&template=bug_report.yml&title=%5BBug%5D%3A+)
· [Suggest a feature](https://github.com/AlejandroSuero/go-insert-locale/issues/new?assignees=&labels=enhancement&projects=&template=feature_request.md&title=%5BFeat%5D%3A+)

**Remember to always follow the [code of conduct](https://github.com/AlejandroSuero/go-insert-locale/blob/main/CODE_OF_CONDUCT.md#contributor-covenant-code-of-conduct)**

</div>

> [!warning]
>
> This project is still in development, so it may not be stable.

## Usage

```bash
go run main.go -input <input-file> -output <output-dir>
```

This will read the JSON file provided in `<input-file>` and insert it into the
directory provided in `<output-dir>`.

```json
{
  "hello": {
    "en": "Hello",
    "es": "Hola"
  }
}
```

This will create the following files in `<output-dir>`:

- `en.json`

```json
{
  "hello": "Hello"
}
```

- `es.json`

```json
{
  "hello": "Hola"
}
```

### Example

```bash
go run main.go -input locales.json -output locales
```

## Contributing

Thank you to everyone that is contributing and to those who want to contribute.
Any contribution is welcomed!

**Quick guide**:

1. [Fork](https://github.com/AlejandroSuero/go-insert-locale/fork) this
   project.
2. Clone your fork (`git clone <fork-URL>`).
3. Add main repo as remote (`git remote add upstream <main-repo-URL>`).
4. Create a branch for your changes (`git switch -c feature/your-feature` or
   `git switch -c fix/your-fix`).
5. Commit your changes (`git commit -m "feat(...): ..."`).
6. Push to your fork (`git push origin <branch-name>`).
7. Open a [PR](https://github.com/AlejandroSuero/go-insert-locale/pulls).

For more information, check
[CONTRIBUTING.md](https://github.com/AlejandroSuero/go-insert-locale/blob/main/CONTRIBUTING.md).

<div align="right">
  (<a href="#readme-top">Back to top</a>)
</div>

### Contributors

Thanks goes to these wonderful people that have contributed to this project:

[![contributors][contributors-img]][contributors-url]

<div align="right">
  (<a href="#readme-top">Back to top</a>)
</div>

[stars-shield]: https://img.shields.io/github/stars/AlejandroSuero/go-insert-locale.svg?style=for-the-badge
[stars-url]: https://github.com/AlejandroSuero/go-insert-locale/stargazers
[issues-shield]: https://img.shields.io/github/issues/AlejandroSuero/go-insert-locale.svg?style=for-the-badge
[issues-url]: https://github.com/AlejandroSuero/go-insert-locale/issues
[contributors-url]: https://github.com/AlejandroSuero/go-insert-locale/graphs/contributors
[contributors-img]: https://contrib.rocks/image?repo=AlejandroSuero/go-insert-locale
