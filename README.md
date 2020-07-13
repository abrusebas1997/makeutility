# MakeUtility project
Community service Certificate generator

## Description
I decided to make a certificate of community service generator because some 
friends from Make School and me are starting a non-profit called “Foodonate”
I thought making a certificate generator would be very helpful for our mission and 
would make the process to create them, way easier. 

### Audience 
Our audience will be any person that wants to enroll in our community service program(mostly students) 
to help people living in homeless shelters to have a healthy and nutritious meal every day. They have 
to complete a certain amount of hours before getting this certificate.


### Installing

1. Project code can be viewed locally by cloning 
2. Then you need to install the pdf package, running
```
$ go get github.com/jung-kurt/gofpdf
```
3. Run the main project and open the pdf file by
```
$ go run main.go && open cert.pdf
```
4. It will ask you for the name of the volunteer and there you go!
![alt text](https://github.com/abrusebas1997/makeutility/blob/master/images/certificate%20(2).jpg)
## Built With

* [Golang](https://golang.org/) - Language 

## Packages used
This project uses the following packages:

* [gopdf](https://github.com/jung-kurt/gofpdf) - to open a pdf file 
* [time](https://golang.org/pkg/time/) - to get exact time when it was generated 
* [sign-svg] (https://willowsystems.github.io/jSignature/#/demo/) - To create sign on svg, easier way to use it in golang

## Authors

* Sebastian Abarca - portfolio can be found at:
https://www.makeschool.com/portfolio/abrusebas16


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
