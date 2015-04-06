package app

import (
	  "github.com/revel/revel";
	  "fmt";
	  "database/sql";
	  "github.com/coopernurse/gorp";
	_ "github.com/denisenkom/go-mssqldb";
)

const (
	TITLE = "DocsVision WebAPI"
)

var Dbm *gorp.DbMap = nil;


type GorpController struct {
	*revel.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() revel.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}


func getParamString(param string, defaultValue string) string {
	p, found := revel.Config.String(param);
	if !found {
		if defaultValue == "" {
			revel.ERROR.Fatal("Cound not find parameter: " + param);
		} else {
			return defaultValue;
		}
	}
	return p;
}

func getConnectionString(dbname_cfg string) string {
	host    := getParamString("db_"+dbname_cfg+".host",             "localhost");
	port    := getParamString("db_"+dbname_cfg+".port",             "1433");
	user    := getParamString("db_"+dbname_cfg+".user",             "sa");
	pass    := getParamString("db_"+dbname_cfg+".password",         "");
	dbname  := getParamString("db_"+dbname_cfg+".name",             dbname_cfg);

	return fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;database=%s", host, port, user, pass, dbname);
}

func init_db(dbname string) *gorp.DbMap {
	connectionString := getConnectionString(dbname);
	var dbm *gorp.DbMap = nil;
	if db, err := sql.Open(getParamString("db_"+dbname+".driver", "mssql"), connectionString); err != nil {
		revel.ERROR.Fatal(err);
	} else {
		dbm = &gorp.DbMap{Db: db};
	}

	return dbm;
}

func init_dbs() {
	Dbm = init_db("docsvision");
}

func init() {

	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)

	revel.OnAppStart(init_dbs);
	revel.InterceptMethod((*GorpController).Begin,          revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit,         revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback,       revel.FINALLY)


        revel.TemplateFuncs["title"] = func(t string) string {
		return TITLE + ": " + revel.Message("ru", "title_"+t)
	}

}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

