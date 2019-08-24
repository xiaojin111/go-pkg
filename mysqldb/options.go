package mysqldb

// Options 是 DbClient 的配置参数
type Options struct {
	// 是否启用日志
	EnableLog bool

	// 连接超时
	// The value must be a decimal number with a unit suffix ("ms", "s", "m", "h"), such as "30s", "0.5m" or "1m30s".
	DialTimeout string

	// 是否转换时间
	ParseTime bool

	// 最大连接数
	MaxConnections int

	// 服务器地址 - host:port
	Address string

	// 用户名
	Username string

	// 密码
	Password string

	// 数据库名
	Database string

	// 字符集
	Charset string

	// 字符排序
	Collation string

	// 区域设置
	Locale string

	// 阻止全局更新
	// 开启选项时，将禁止没有 WHERE 语句的 DELETE 或 UPDATE 操作执行，否则抛出 error
	BlockGlobalUpdate bool
}

// Option 是设置 Options 的函数
type Option func(*Options)

// DialTimeout 设置连接超时时间
// The value must be a decimal number with a unit suffix ("ms", "s", "m", "h"), such as "30s", "0.5m" or "1m30s".
func DialTimeout(timeout string) Option {
	return func(o *Options) {
		o.DialTimeout = timeout
	}
}

// Address 设置服务器地址 - host:port
func Address(addr string) Option {
	return func(o *Options) {
		o.Address = addr
	}
}

// Username 设置用户名
func Username(username string) Option {
	return func(o *Options) {
		o.Username = username
	}
}

// Password 设置密码
func Password(pwd string) Option {
	return func(o *Options) {
		o.Password = pwd
	}
}

// Database 设置数据库名
func Database(db string) Option {
	return func(o *Options) {
		o.Database = db
	}
}

// EnableLog 设置是否启用日志
func EnableLog(enable bool) Option {
	return func(o *Options) {
		o.EnableLog = enable
	}
}

// MaxConnections 设置最大连接数
func MaxConnections(maxConns int) Option {
	return func(o *Options) {
		o.MaxConnections = maxConns
	}
}

// Charset 设置字符集
func Charset(charset string) Option {
	return func(o *Options) {
		o.Charset = charset
	}
}

// Collation 设置字符集排序
func Collation(collation string) Option {
	return func(o *Options) {
		o.Collation = collation
	}
}

// ParseTime 设置转换时间
func ParseTime(parseTime bool) Option {
	return func(o *Options) {
		o.ParseTime = parseTime
	}
}

// Locale 设置区域设置
func Locale(locale string) Option {
	return func(o *Options) {
		o.Locale = locale
	}
}

// BlockGlobalUpdate 设置是否阻止全局更新
func BlockGlobalUpdate(block bool) Option {
	return func(o *Options) {
		o.BlockGlobalUpdate = block
	}
}

// DefaultOptions 返回默认配置的 Options
func DefaultOptions() Options {
	return Options{
		Address:           "localhost:3306",
		DialTimeout:       "10s", // 默认连接超时时间为10秒
		EnableLog:         false,
		MaxConnections:    0,
		Charset:           "utf8mb4",
		Collation:         "utf8mb4_general_ci",
		ParseTime:         true,
		Locale:            "UTC", // 注意: 这里字母必须大写，否则找不到 Timezone 文件
		BlockGlobalUpdate: true,
	}
}

// NewOptions 设置 Options
func NewOptions(opt ...Option) Options {
	opts := DefaultOptions()

	for _, o := range opt {
		o(&opts)
	}

	return opts
}
