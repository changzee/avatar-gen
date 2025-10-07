# Golang 卡通头像生成器 (SVG)

这是一个使用 Go 语言编写的卡通头像生成器。它可以根据用户指定的国家和性别，通过分层合成的方式，随机生成独一无二的 **SVG** 格式的卡通头像。

## ✨ 功能特性

-   **高度可定制**：通过替换或添加 **SVG** 格式的图片素材，轻松定制不同风格的头像。
-   **矢量图形**：生成的 SVG 头像可以无限缩放而不失真。
-   **灵活的素材系统**：支持按国家、性别进行素材分类，并提供通用素材作为备用。
-   **随机生成**：每次运行都会生成一个不重复的头像。
-   **易于扩展**：代码结构清晰，方便进行二次开发。

## 📂 项目结构

```
.
├── assets/                # SVG 素材目录
│   ├── china/             # 国家：中国
│   │   ├── male/          # 性别：男性
│   │   │   ├── 02_clothes/
│   │   │   └── 05_hair/
│   │   └── ...
│   ├── usa/               # 国家：美国
│   │   └── ...
│   └── common/            # 通用素材目录
│       ├── common/        # 男女通用
│       │   ├── 00_background/
│       │   └── 04_eyes/
│       └── male/          # 男性通用
│           ├── 01_body/
│           └── 03_mouth/
├── cmd/
│   └── avatar-generator/
│       └── main.go        # 程序主入口
├── go.mod                 # Go 模块文件
├── go.sum
├── output/                # 生成的头像输出目录
├── pkg/
│   └── avatar/
│       └── generator.go   # 核心头像生成器逻辑
└── README.md              # 项目说明
```

## 🚀 如何开始

### 1. 前提条件

-   安装 [Go](https://golang.org/dl/) (版本 1.16 或更高)。

### 2. 构建项目

你可以直接运行，Go 会自动处理构建：

```bash
go build -o avatar-generator ./cmd/avatar-generator
```

### 3. 运行程序

通过命令行参数指定国家 (`--country`) 和性别 (`--gender`) 来生成头像。

**示例：**

```bash
# 生成一个中国男性的 SVG 头像
go run ./cmd/avatar-generator --country=china --gender=male

# 生成一个美国女性的 SVG 头像，并指定输出目录
go run ./cmd/avatar-generator --country=usa --gender=female --output=./my_avatars
```

生成的头像将默认保存在 `output/` 目录下，并带有一个包含时间戳的唯一文件名，格式为 `.svg`。

## 🎨 如何添加或修改素材

这是本项目的核心乐趣所在！

1.  **准备素材**：
    *   创建你的 `.svg` 格式的矢量图素材。你可以使用任何矢量图形编辑器（如 Inkscape, Adobe Illustrator）。
    *   **重要**：为了方便组合，建议所有 SVG 文件都使用相同的 `viewBox` 属性（例如 `viewBox="0 0 512 512"`）。这样可以确保所有图层内的坐标系一致。
    *   在每个SVG文件中，只包含该图层需要的图形元素，不要包含 `<svg>` 标签外的其他内容。

2.  **组织素材**：
    *   图层文件夹必须以 `数字_名称` 的格式命名（如 `00_background`, `01_body`），数字决定了它们的叠加顺序。
    *   将素材放入对应的 `assets/{国家}/{性别}/{图层}` 目录中。
    *   如果某个素材是通用的（不限国家或性别），可以将其放入 `assets/common/...` 下的相应目录中。程序在找不到特定国家的素材时，会自动使用通用素材。

**素材查找优先级:**

程序会按照以下顺序寻找素材，一旦找到，便会停止搜索该图层：
1.  `assets/{country}/{gender}/`
2.  `assets/common/{gender}/`
3.  `assets/{country}/common/`
4.  `assets/common/common/`

祝你玩得开心，创造出属于你自己的独特 SVG 头像！