// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package driver

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"go.etcd.io/etcd/clientv3"
)

// Etcd driver
type Etcd struct {
	Client *clientv3.Client
}

// NewEtcdDriver create a new instance
func NewEtcdDriver() (*Etcd, error) {
	result := &Etcd{}

	result.Client, err = clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(viper.GetString("database.etcd.endpoints"), ","),
		DialTimeout: viper.GetInt("database.etcd.timeout") * time.Second,
		Username:    viper.GetString("database.etcd.username"),
		Password:    viper.GetString("database.etcd.password"),
	})

	if err != nil {
		return result, err
	}

	return result, nil
}

// Put sets a record
func (e *Etcd) Put(key, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetInt("database.etcd.timeout")*time.Second)

	_, err := e.Client.Put(ctx, key, value)

	defer cancel()

	if err != nil {
		return err
	}

	return nil
}

// PutWithLease sets a record
func (e *Etcd) PutWithLease(key, value string, leaseId clientv3.LeaseID) error {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetInt("database.etcd.timeout")*time.Second)

	_, err := e.Client.Put(ctx, key, value, clientv3.WithLease(leaseId))

	defer cancel()

	if err != nil {
		return err
	}

	return nil
}

// Get gets a record value
func (e *Etcd) Get(key string) (map[string]string, error) {
	result := make(map[string]string, 0)

	ctx, cancel := context.WithTimeout(context.Background(), viper.GetInt("database.etcd.timeout")*time.Second)

	resp, err := e.Client.Get(ctx, key, clientv3.WithPrefix())

	defer cancel()

	if err != nil {
		return result, err
	}

	for _, ev := range resp.Kvs {
		result[string(ev.Key)] = string(ev.Value)
	}

	return result, nil
}

// Delete deletes a record
func (e *Etcd) Delete(key string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetInt("database.etcd.timeout")*time.Second)

	dresp, err := e.Client.Delete(ctx, key, clientv3.WithPrefix())

	defer cancel()

	if err != nil {
		return 0, err
	}

	return dresp.Deleted, nil
}

// CreateLease creates a lease
func (e *Etcd) CreateLease(seconds int64) (clientv3.LeaseID, error) {
	var result clientv3.LeaseID

	ctx, cancel := context.WithTimeout(context.Background(), viper.GetInt("database.etcd.timeout")*time.Second)

	resp, err := e.Client.Grant(ctx, seconds)

	defer cancel()

	if err != nil {
		return result, err
	}

	return resp.ID, nil
}

// RenewLease renews a lease
func (e *Etcd) RenewLease(leaseId clientv3.LeaseID) error {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetInt("database.etcd.timeout")*time.Second)

	_, err := e.Client.KeepAliveOnce(ctx, leaseId)

	defer cancel()

	if err != nil {
		return err
	}

	return nil
}

// Close closes the etcd connection
func (e *Etcd) Close() {
	e.Client.Close()
}
