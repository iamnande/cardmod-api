// Code generated by entc, DO NOT EDIT.

package database

import (
	"context"
	"fmt"
	"log"

	"github.com/iamnande/cardmod/internal/database/migrate"

	"github.com/iamnande/cardmod/internal/database/card"
	"github.com/iamnande/cardmod/internal/database/item"
	"github.com/iamnande/cardmod/internal/database/limitbreak"
	"github.com/iamnande/cardmod/internal/database/magic"
	"github.com/iamnande/cardmod/internal/database/refinement"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Card is the client for interacting with the Card builders.
	Card *CardClient
	// Item is the client for interacting with the Item builders.
	Item *ItemClient
	// LimitBreak is the client for interacting with the LimitBreak builders.
	LimitBreak *LimitBreakClient
	// Magic is the client for interacting with the Magic builders.
	Magic *MagicClient
	// Refinement is the client for interacting with the Refinement builders.
	Refinement *RefinementClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Card = NewCardClient(c.config)
	c.Item = NewItemClient(c.config)
	c.LimitBreak = NewLimitBreakClient(c.config)
	c.Magic = NewMagicClient(c.config)
	c.Refinement = NewRefinementClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("database: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("database: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Card:       NewCardClient(cfg),
		Item:       NewItemClient(cfg),
		LimitBreak: NewLimitBreakClient(cfg),
		Magic:      NewMagicClient(cfg),
		Refinement: NewRefinementClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:     cfg,
		Card:       NewCardClient(cfg),
		Item:       NewItemClient(cfg),
		LimitBreak: NewLimitBreakClient(cfg),
		Magic:      NewMagicClient(cfg),
		Refinement: NewRefinementClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Card.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Card.Use(hooks...)
	c.Item.Use(hooks...)
	c.LimitBreak.Use(hooks...)
	c.Magic.Use(hooks...)
	c.Refinement.Use(hooks...)
}

// CardClient is a client for the Card schema.
type CardClient struct {
	config
}

// NewCardClient returns a client for the Card from the given config.
func NewCardClient(c config) *CardClient {
	return &CardClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `card.Hooks(f(g(h())))`.
func (c *CardClient) Use(hooks ...Hook) {
	c.hooks.Card = append(c.hooks.Card, hooks...)
}

// Create returns a create builder for Card.
func (c *CardClient) Create() *CardCreate {
	mutation := newCardMutation(c.config, OpCreate)
	return &CardCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Card entities.
func (c *CardClient) CreateBulk(builders ...*CardCreate) *CardCreateBulk {
	return &CardCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Card.
func (c *CardClient) Update() *CardUpdate {
	mutation := newCardMutation(c.config, OpUpdate)
	return &CardUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CardClient) UpdateOne(ca *Card) *CardUpdateOne {
	mutation := newCardMutation(c.config, OpUpdateOne, withCard(ca))
	return &CardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CardClient) UpdateOneID(id int) *CardUpdateOne {
	mutation := newCardMutation(c.config, OpUpdateOne, withCardID(id))
	return &CardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Card.
func (c *CardClient) Delete() *CardDelete {
	mutation := newCardMutation(c.config, OpDelete)
	return &CardDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CardClient) DeleteOne(ca *Card) *CardDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CardClient) DeleteOneID(id int) *CardDeleteOne {
	builder := c.Delete().Where(card.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CardDeleteOne{builder}
}

// Query returns a query builder for Card.
func (c *CardClient) Query() *CardQuery {
	return &CardQuery{
		config: c.config,
	}
}

// Get returns a Card entity by its id.
func (c *CardClient) Get(ctx context.Context, id int) (*Card, error) {
	return c.Query().Where(card.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CardClient) GetX(ctx context.Context, id int) *Card {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CardClient) Hooks() []Hook {
	return c.hooks.Card
}

// ItemClient is a client for the Item schema.
type ItemClient struct {
	config
}

// NewItemClient returns a client for the Item from the given config.
func NewItemClient(c config) *ItemClient {
	return &ItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `item.Hooks(f(g(h())))`.
func (c *ItemClient) Use(hooks ...Hook) {
	c.hooks.Item = append(c.hooks.Item, hooks...)
}

// Create returns a create builder for Item.
func (c *ItemClient) Create() *ItemCreate {
	mutation := newItemMutation(c.config, OpCreate)
	return &ItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Item entities.
func (c *ItemClient) CreateBulk(builders ...*ItemCreate) *ItemCreateBulk {
	return &ItemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Item.
func (c *ItemClient) Update() *ItemUpdate {
	mutation := newItemMutation(c.config, OpUpdate)
	return &ItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ItemClient) UpdateOne(i *Item) *ItemUpdateOne {
	mutation := newItemMutation(c.config, OpUpdateOne, withItem(i))
	return &ItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ItemClient) UpdateOneID(id int) *ItemUpdateOne {
	mutation := newItemMutation(c.config, OpUpdateOne, withItemID(id))
	return &ItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Item.
func (c *ItemClient) Delete() *ItemDelete {
	mutation := newItemMutation(c.config, OpDelete)
	return &ItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ItemClient) DeleteOne(i *Item) *ItemDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ItemClient) DeleteOneID(id int) *ItemDeleteOne {
	builder := c.Delete().Where(item.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ItemDeleteOne{builder}
}

// Query returns a query builder for Item.
func (c *ItemClient) Query() *ItemQuery {
	return &ItemQuery{
		config: c.config,
	}
}

// Get returns a Item entity by its id.
func (c *ItemClient) Get(ctx context.Context, id int) (*Item, error) {
	return c.Query().Where(item.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ItemClient) GetX(ctx context.Context, id int) *Item {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ItemClient) Hooks() []Hook {
	return c.hooks.Item
}

// LimitBreakClient is a client for the LimitBreak schema.
type LimitBreakClient struct {
	config
}

// NewLimitBreakClient returns a client for the LimitBreak from the given config.
func NewLimitBreakClient(c config) *LimitBreakClient {
	return &LimitBreakClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `limitbreak.Hooks(f(g(h())))`.
func (c *LimitBreakClient) Use(hooks ...Hook) {
	c.hooks.LimitBreak = append(c.hooks.LimitBreak, hooks...)
}

// Create returns a create builder for LimitBreak.
func (c *LimitBreakClient) Create() *LimitBreakCreate {
	mutation := newLimitBreakMutation(c.config, OpCreate)
	return &LimitBreakCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of LimitBreak entities.
func (c *LimitBreakClient) CreateBulk(builders ...*LimitBreakCreate) *LimitBreakCreateBulk {
	return &LimitBreakCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for LimitBreak.
func (c *LimitBreakClient) Update() *LimitBreakUpdate {
	mutation := newLimitBreakMutation(c.config, OpUpdate)
	return &LimitBreakUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LimitBreakClient) UpdateOne(lb *LimitBreak) *LimitBreakUpdateOne {
	mutation := newLimitBreakMutation(c.config, OpUpdateOne, withLimitBreak(lb))
	return &LimitBreakUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LimitBreakClient) UpdateOneID(id int) *LimitBreakUpdateOne {
	mutation := newLimitBreakMutation(c.config, OpUpdateOne, withLimitBreakID(id))
	return &LimitBreakUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for LimitBreak.
func (c *LimitBreakClient) Delete() *LimitBreakDelete {
	mutation := newLimitBreakMutation(c.config, OpDelete)
	return &LimitBreakDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *LimitBreakClient) DeleteOne(lb *LimitBreak) *LimitBreakDeleteOne {
	return c.DeleteOneID(lb.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *LimitBreakClient) DeleteOneID(id int) *LimitBreakDeleteOne {
	builder := c.Delete().Where(limitbreak.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LimitBreakDeleteOne{builder}
}

// Query returns a query builder for LimitBreak.
func (c *LimitBreakClient) Query() *LimitBreakQuery {
	return &LimitBreakQuery{
		config: c.config,
	}
}

// Get returns a LimitBreak entity by its id.
func (c *LimitBreakClient) Get(ctx context.Context, id int) (*LimitBreak, error) {
	return c.Query().Where(limitbreak.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LimitBreakClient) GetX(ctx context.Context, id int) *LimitBreak {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *LimitBreakClient) Hooks() []Hook {
	return c.hooks.LimitBreak
}

// MagicClient is a client for the Magic schema.
type MagicClient struct {
	config
}

// NewMagicClient returns a client for the Magic from the given config.
func NewMagicClient(c config) *MagicClient {
	return &MagicClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `magic.Hooks(f(g(h())))`.
func (c *MagicClient) Use(hooks ...Hook) {
	c.hooks.Magic = append(c.hooks.Magic, hooks...)
}

// Create returns a create builder for Magic.
func (c *MagicClient) Create() *MagicCreate {
	mutation := newMagicMutation(c.config, OpCreate)
	return &MagicCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Magic entities.
func (c *MagicClient) CreateBulk(builders ...*MagicCreate) *MagicCreateBulk {
	return &MagicCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Magic.
func (c *MagicClient) Update() *MagicUpdate {
	mutation := newMagicMutation(c.config, OpUpdate)
	return &MagicUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MagicClient) UpdateOne(m *Magic) *MagicUpdateOne {
	mutation := newMagicMutation(c.config, OpUpdateOne, withMagic(m))
	return &MagicUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MagicClient) UpdateOneID(id int) *MagicUpdateOne {
	mutation := newMagicMutation(c.config, OpUpdateOne, withMagicID(id))
	return &MagicUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Magic.
func (c *MagicClient) Delete() *MagicDelete {
	mutation := newMagicMutation(c.config, OpDelete)
	return &MagicDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MagicClient) DeleteOne(m *Magic) *MagicDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MagicClient) DeleteOneID(id int) *MagicDeleteOne {
	builder := c.Delete().Where(magic.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MagicDeleteOne{builder}
}

// Query returns a query builder for Magic.
func (c *MagicClient) Query() *MagicQuery {
	return &MagicQuery{
		config: c.config,
	}
}

// Get returns a Magic entity by its id.
func (c *MagicClient) Get(ctx context.Context, id int) (*Magic, error) {
	return c.Query().Where(magic.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MagicClient) GetX(ctx context.Context, id int) *Magic {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MagicClient) Hooks() []Hook {
	return c.hooks.Magic
}

// RefinementClient is a client for the Refinement schema.
type RefinementClient struct {
	config
}

// NewRefinementClient returns a client for the Refinement from the given config.
func NewRefinementClient(c config) *RefinementClient {
	return &RefinementClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `refinement.Hooks(f(g(h())))`.
func (c *RefinementClient) Use(hooks ...Hook) {
	c.hooks.Refinement = append(c.hooks.Refinement, hooks...)
}

// Create returns a create builder for Refinement.
func (c *RefinementClient) Create() *RefinementCreate {
	mutation := newRefinementMutation(c.config, OpCreate)
	return &RefinementCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Refinement entities.
func (c *RefinementClient) CreateBulk(builders ...*RefinementCreate) *RefinementCreateBulk {
	return &RefinementCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Refinement.
func (c *RefinementClient) Update() *RefinementUpdate {
	mutation := newRefinementMutation(c.config, OpUpdate)
	return &RefinementUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RefinementClient) UpdateOne(r *Refinement) *RefinementUpdateOne {
	mutation := newRefinementMutation(c.config, OpUpdateOne, withRefinement(r))
	return &RefinementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RefinementClient) UpdateOneID(id int) *RefinementUpdateOne {
	mutation := newRefinementMutation(c.config, OpUpdateOne, withRefinementID(id))
	return &RefinementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Refinement.
func (c *RefinementClient) Delete() *RefinementDelete {
	mutation := newRefinementMutation(c.config, OpDelete)
	return &RefinementDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RefinementClient) DeleteOne(r *Refinement) *RefinementDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RefinementClient) DeleteOneID(id int) *RefinementDeleteOne {
	builder := c.Delete().Where(refinement.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RefinementDeleteOne{builder}
}

// Query returns a query builder for Refinement.
func (c *RefinementClient) Query() *RefinementQuery {
	return &RefinementQuery{
		config: c.config,
	}
}

// Get returns a Refinement entity by its id.
func (c *RefinementClient) Get(ctx context.Context, id int) (*Refinement, error) {
	return c.Query().Where(refinement.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RefinementClient) GetX(ctx context.Context, id int) *Refinement {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *RefinementClient) Hooks() []Hook {
	return c.hooks.Refinement
}
