package controllers

import (
	"encoding/json"
	"revel-demo/app/models"

	"github.com/revel/revel"
)

type UserCtrl struct {
	GorpController
}

func (c UserCtrl) parseUser() (models.User, error) {
	user := models.User{}
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	return user, err
}

func (c UserCtrl) Add() revel.Result {
	if user, err := c.parseUser(); err != nil {
		return c.RenderText("Unable to parse the User from JSON.")
	} else {
		// Validate the model
		user.Validate(c.Validation)
		if c.Validation.HasErrors() {
			// Do something better here!
			return c.RenderText("You have error with the User.")
		} else {
			if err := c.Txn.Insert(&user); err != nil {
				return c.RenderText(
					"Error inserting record into database!")
			} else {
				return c.RenderJson(user)
			}
		}
	}
}

func (c UserCtrl) Get(id int64) revel.Result {
	user := new(models.User)
	err := c.Txn.SelectOne(user,
		`SELECT * FROM User WHERE id = ?`, id)
	if err != nil {
		return c.RenderText("Error.  User probably doesn't exist.")
	}
	return c.RenderJson(user)
}

func (c UserCtrl) List() revel.Result {
	lastId := parseIntOrDefault(c.Params.Get("lid"), -1)
	limit := parseUintOrDefault(c.Params.Get("limit"), uint64(25))
	users, err := c.Txn.Select(models.User{},
		`SELECT * FROM User WHERE id > ? LIMIT ?`, lastId, limit)
	if err != nil {
		return c.RenderText(
			"Error trying to get records from DB.")
	}
	return c.RenderJson(users)
}

func (c UserCtrl) Update(id int64) revel.Result {
	user, err := c.parseUser()
	if err != nil {
		return c.RenderText("Unable to parse the User from JSON.")
	}
	// Ensure the Id is set.
	user.Id = id
	success, err := c.Txn.Update(&user)
	if err != nil || success == 0 {
		return c.RenderText("Unable to update user.")
	}
	return c.RenderText("Updated %v", id)
}

func (c UserCtrl) Delete(id int64) revel.Result {
	success, err := c.Txn.Delete(&models.User{Id: id})
	if err != nil || success == 0 {
		return c.RenderText("Failed to remove User")
	}
	return c.RenderText("Deleted %v", id)
}
