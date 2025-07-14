package bootstrap

import (
	"context"
	"fmt"
	"time"
)

type Application struct {
	Env *Env
	Db  *DB
}

func App() *Application {
	app := &Application{}

	app.Env = NewEnv()
	var err error
	if app.Db, err = NewDb(app.Env); err != nil {
		panic(fmt.Sprint("fail to init db, err = ", err))
	}

	return app
}

func (app *Application) Close() error {
	fmt.Println("app is destoring")
	if app.Db == nil {
		return nil
	}
	sqlDb, err := app.Db.Db.DB()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- sqlDb.Close()
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return fmt.Errorf("数据库关闭超时: %v", ctx.Err())
	}
}
