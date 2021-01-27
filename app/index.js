function FormatPrimeDetail(value){
  if (value.is_prime) {
    return 'Prime Number'
  } else {
    return 'Composite Number with factors ' + FormatArray(value.factors)
  }
}

function FormatArray(numbers){
  return '[' + numbers.join(', ') + ']'
}

// duration in Nanosecond
function FormatTime(duration){
  micro = Math.floor(duration/1000)
  milli = micro/1000
  return milli
}

function FormatStrToArray(str){
  return str.replace(/\s/g, '').split(',').map(x=> +x).filter(x=>x!=0 && x < 10000000)
}

function RenderData(data, targetNode, isGoroutine){
  targetNode.html('')
  var $table = $('<table style="width:100%">')
  var html = []

  html.push(`<caption>Prime Factors ${isGoroutine ? 'with' : 'without'} Goroutine</caption>`)
  html.push('<tr><th>Prime</th><th>Factors</th></tr>')
  $.each(data.results, function( index, value ) {
    html.push('<tr><td>'+ value.digit +'</td><td>'+ FormatPrimeDetail(value) +'</td></tr>')
  });

  $table.html(html.join())
  $note = $(`<div>numbers: ${FormatArray(data.numbers)} </div><div> duration time: ${FormatTime(data.duration)}ms </div>`)

  targetNode.append($table, $note)
}

function GetData(data){
  return $.ajax({
    url: '/prime-factors',
    method: 'post',
    data:JSON.stringify(data),
    contentType: 'application/json',
    async: false,
  })
}

// return duration time
function PrimeFactorsRender(example_type, numbers, is_goroutine){
  var res = GetData({
    example_type: example_type,
    is_goroutine: is_goroutine,
    numbers:numbers
  })

  var $node = is_goroutine ? $('#main-right') : $('#main-left')
  var data = res.responseJSON.data

  RenderData(data, $node, is_goroutine)
  return data.duration
}

function DifferentTimeRender(time, targetNode){
  targetNode.html(`Different Time: ${FormatTime(time)}ms (${time > 0 ? 'faster': 'slower'})`)
}

$('.control button#custom').click(function(e){
  console.log("custom")
  $('input#numbers').show()
  $('#numbers').focus()
})
$('.control button').not('#custom').click(function(e){
  console.log(e.target.id)
  example_type = e.target.id
  Init(example_type, [])
})
$(document).click(function(e){
  if(!['numbers', 'custom'].includes(e.target.id)){
    $('input#numbers').hide()
  }
})
$('input#numbers').keypress(function(e){
  if (!e.key.match(/[0-9]|,/)){
    e.preventDefault();
  }
  if (e.key == 'Enter'){
    value = e.target.value
    e.target.value = ''
    Init("custom", FormatStrToArray(value))
  }
})

function Init(example_type, numbers){
  dWithout = PrimeFactorsRender(example_type, numbers, false)
  dWith = PrimeFactorsRender(example_type, numbers, true)
  DifferentTimeRender(dWithout - dWith, $('#diff-time'))
  console.log('done')
}

Init('medium', [])