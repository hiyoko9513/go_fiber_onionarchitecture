# Release Drafter
[参照](https://zenn.dev/kounoike/articles/20220211-easy-generate-release-notes)

## 手順
1. ブランチの作成(auto labeler利用)
    - feature/(:機能追加、改善) → feature
    - fix/(:バグ修正) → bug
    - chore/(:メンテナンス) → chore
    - refactor/(:リファクタリング) → refactor
    - doc/(:ドキュメント) → documentation
2. PR作成(developに向け)
    - prのタイトルがリリースノートに記載される
3. marge(release noteに自動追加される)

## バージョンについて
1. major
    - prのlabelに[major]
2. minor
    - prのlabelに[minor]
3. patch
    - prのlabelに[patch]

prに上記のラベルを付与した場合、バージョンがアップ出来る。
defaultはpatchになっている。

## まとめ
- ブランチ名に注意する(label用)
- PRのタイトルに注意する(release note記載文言)
- バージョンを上げたい場合PRに直接書き込む
