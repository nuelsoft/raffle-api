package static

const (
	AlreadyWon = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Already Won</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
          integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">

</head>
<body>
<div class="card">
    <div class="card-header font-weight-bolder">
        WINNER
    </div>
    <div class="card-body container-fluid">
        <h5 class="card-title">You got the magic!</h5>
		<p class="card-text">You won again. Keep winning and don't forgot to share the link 
		https://badcommentsmoviepromo.com with your friends and family</p>
        <a href="https://badcommentsmoviepromo.com" class="btn btn-primary">See Other winners</a>
    </div>
</div>
</body>
</html>
`

	Congrats = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Congratulations</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
          integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
</head>
<body>
<div class="card">
    <div class="card-header font-weight-bolder">
        Congratulations
    </div>
    <div class="card-body container-fluid">
        <h5 class="card-title">You WON</h5>
		<p class="card-text">
		Thank you for participating in the "BAD COMMENTS THE MOVIE" Promo. 
		Congratulations! You have been selected as one of the random winners for today's draw. 
		Your win grants you free access to attend  "BAD COMMENTS THE MOVIE PREMIERE" this Easter 
		and a photoshoot with your favourite "BAD COMMENTS MOVIE STARS" like Jim Iyke, Ini Edo, Timaya, 
		Daddy Freeze etc. Wait😊 it has not finished yet🕺🏾. You will also receive free roundtrip transportation 
		from your location to the "BAD COMMENTS THE MOVIE" Premiere venue and a V.I.P red carpet treatment. 
		Oh wait! This is not the end, You can continue to play on for as many times as you wish to be among
		the 10 lucky winners to win grand prizes of 1k USD each and all expense paid trip to Zanzibar with
		Jim Iyke and Ini Edo for 3days starting from Easter Sunday. It only gets better with 
		<b>"BAD COMMENTS THE MOVIE"</b>. P.S send the BAD COMMENTS MOVIE PROMO LINK 
		<a href="https://badcommentsmoviepromo.com">https://badcommentsmoviepromo.com</a> to your 
		family and friends for them to win too. Yes oo! All we do is WIN! WIN! WIN!💪🏾🕺🏾
		</p>
        <a href="https://badcommentsmoviepromo.com" class="btn btn-primary">See Other winners</a>
    </div>
</div>
</body>
</html>
`

	InvalidReq = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Invalid Request</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
          integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">

</head>
<body>
<div class="card">
    <div class="card-header font-weight-bolder">
        INVALID REQUEST
    </div>
    <div class="card-body container-fluid">
        <h5 class="card-title">Don't do that again</h5>
        <p class="card-text">Request seems to be poorly formatted or missing some query parameters</p>
        <a href="https://badcommentsmoviepromo.com" class="btn btn-primary">Place raffle</a>
    </div>
</div>
</body>
</html>
`

	InvalidTx = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Invalid Transaction</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
          integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">

</head>
<body>
<div class="card">
    <div class="card-header font-weight-bolder">
        TRANSACTION ERROR
    </div>
    <div class="card-body container-fluid">
        <h5 class="card-title">Couldn't verify Transaction</h5>
        <p class="card-text">Please proceed to make raffle and payment.</p>
        <a href="https://badcommentsmoviepromo.com" class="btn btn-primary">Go to Raffle Page</a>
    </div>
</div>
</body>
</html>
`

	PaymentRef = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Used Payment Ref</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
          integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">

</head>
<body>
<div class="card">
    <div class="card-header font-weight-bolder">
        PAYMENT ERROR
    </div>
    <div class="card-body container-fluid">
        <h5 class="card-title">REF has already been used</h5>
        <p class="card-text">Please proceed to make another raffle and payment.</p>
        <a href="https://badcommentsmoviepromo.com" class="btn btn-primary">Go to Raffle Page</a>
    </div>
</div>
</body>
</html>`

	Sorry = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Sorry</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
          integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">

</head>
<body>
<div class="card">
    <div class="card-header font-weight-bolder">
        Try again
    </div>
    <div class="card-body container-fluid">
        <h5 class="card-title">Oops!</h5>
		<p class="card-text">Oops! Try again to be amongst the winners for today's draw.
		 Your win grants you free access to attend  "BAD COMMENTS THE MOVIE PREMIERE" this Easter and a 
		 photoshoot with your favourite "BAD COMMENTS MOVIE STARS" like Jim Iyke, Ini Edo, Timaya, Daddy Freeze etc. 
		 Wait😊 it has not finished yet🕺🏾. You will also receive free roundtrip transportation from your location 
		 to the "BAD COMMENTS THE MOVIE" Premiere venue and a V.I.P red carpet treatment. Oh wait! This is not 
		 the end, You can continue to play on for as many times as you wish to be among the 10 lucky winners to 
		 win grand prizes of 1k USD each and all expense paid trip to Zanzibar with Jim Iyke and Ini Edo for 3 days 
		 starting from Easter Sunday. It only gets better with "BAD COMMENTS THE MOVIE". P.S. Send the BAD COMMENTS MOVIE
		  PROMO LINK https://badcommentsmoviepromo.com/ to your family and friends for them to win too. 
		Yes oo! Play again & don't give up bcos All we do is WIN! WIN! WIN! No matter what.💪🏾💪🏾</p>
        <a href="https://badcommentsmoviepromo.com" class="btn btn-primary">Try again</a>
        <a href="https://api-badcommentsmoviepromo.herokuapp.com/winners" class="btn btn-secondary ">See winners</a>

    </div>
</div>
</body>
</html>`

	Winners = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Winners</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
          integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">

</head>
<body>
<div class="container pt-3">
<div class="float-right">
        <a  href="https://badcommentsmoviepromo.com">Go Back to Raffle</a>
    </div>
    <h1>Cheers to our <span class="text-info font-weight-bolder">Winners</span></h1>
    <div class="card mt-3">
        <table class="table table-striped">
            <thead>
            <tr>
                <th scope="col">Name</th>
                <th scope="col">Phone</th>
                <th scope="col">Email</th>
            </tr>
            </thead>
            <tbody>
            %s
            </tbody>
        </table>

    </div>

</div>
</body>
</html>`
)
