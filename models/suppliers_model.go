package models

import (
	"fmt"

	"net/http"
	cm "pnp/echo-rest/common"
	"pnp/echo-rest/db"

	"github.com/labstack/echo"
	ex "github.com/wolvex/go/error"
)

//var errMessage string
//var errs *ex.AppError
var supp cm.Suppliers
var suppObj []cm.Suppliers

func FetchSuppliers() (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	con := db.CreateCon()

	sqlQuery := `SELECT
					IFNULL(CompanyName,'')CompanyName,
					IFNULL(ContactName,'') ContactName,
					IFNULL(ContactTitle,'') ContactTitle,
					IFNULL(Address,'') Address,
					IFNULL(City,'') City,
					IFNULL(PostalCode,'') PostalCode,
					IFNULL(Country,'') Country,
					IFNULL(Phone,'') Phone ,
					IFNULL(Fax,'') Fax,
					IFNULL(HomePage,'') HomePage
				FROM suppliers `

	rows, err := con.Query(sqlQuery)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&supp.CompanyName, &supp.ContactName, &supp.ContactTitle,
			&supp.Address, &supp.City, &supp.PostalCode,
			&supp.Country, &cust.Phone, &supp.Fax,
			&supp.HomePage)

		if err != nil {
			errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
			errMessage = err.Error()
			return res, err
		}

		suppObj = append(suppObj, supp)

	}

	res.Status = http.StatusOK
	res.Message = "succses"
	res.Data = suppObj

	return res, nil
}

//StoreSuppliers ...
func StoreSuppliers(e echo.Context) (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	req := new(cm.Suppliers)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `INSERT INTO suppliers (CompanyName,ContactName,ContactTitle,Address,City,PostalCode,Country,Phone,Fax,HomePage)
					 VALUES (?,?,?,?,?,?,?,?,?,?)`

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	result, err := stmt.Exec(req.CompanyName, req.ContactName, req.ContactTitle, req.Address,
		req.City, req.PostalCode, req.Country, req.Phone, req.Fax, req.HomePage)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "succes"

	// res.Data = map[int]int{
	// 	"SupplierID ADD": req.SupplierID,
	// }

	return res, nil
}

//UpdateSuppliers ...
func UpdateSuppliers(e echo.Context) (res Response, err error) {

	req := new(cm.Suppliers)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `UPDATE suppliers SET CompanyName = ?, ContactName = ?, ContactTitle = ? WHERE  SupplierID = ?`

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	result, err := stmt.Exec(req.CompanyName, req.ContactName, req.ContactTitle, req.SupplierID)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "succes"

	// res.Data = map[int]int{
	// 	"row affected :": req.SupplierID,
	// }

	return res, nil
}

//DeleteCustomer ...
func DeleteSuppliers(e echo.Context) (res Response, err error) {

	req := new(cm.Suppliers)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `DELETE FROM suppliers WHERE  SupplierID = ?`

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	result, err := stmt.Exec(req.SupplierID)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "succes"

	// res.Data = map[string]string{
	// 	"row deleted :": req.CustomerID,
	// }

	return res, nil
}
