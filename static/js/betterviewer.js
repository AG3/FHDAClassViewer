var subjects_dict = []
var course_dict = []

var stat_sub = false
var stat_crse = false

var curTable = []
var curSize = 0

//---------------------------------

function buildRow(obj){
    raw_row = '<tr class="active" onclick="addClassToView('+curSize+')"><td>'+
    obj['title']+'</td><td>'+
    obj['days']+'</td><td>'+
    obj['start_time']+'</td><td>'+
    obj['end_time']+'</td><td>'+
    obj['instructor']+'</td><td>'
    if(obj['online_content']=='1'){
        raw_row+='Yes'
    }else{
        raw_row+='No'
    }
    raw_row+='</td></tr>'
    curSize++
    curTable.push(obj)

    $('tbody').append(raw_row)
    //AddClass(obj)
}

function requestClass(sub, crse){
    $.ajax({
        type:"GET",
        url:"/class",
        data:{
            "subject":sub,
            "course":crse
        }
    }).done(function(msg){
        classes = msg.split('|')
        for (i in classes){
            buildRow(JSON.parse(classes[i]))
        }
    })
}

function setText(e, tar){
    $('tbody').empty()
    if(tar == 'sub'){
        $('#subject_input').val(e.text)
        $('#course_input').val('')
        validateSub()
        return
    }
    $('#course_input').val(e.text)
    validateCrse()

    if(stat_sub && stat_crse){
        ClassEleSet.remove()
        requestClass($('#subject_input').val(),$('#course_input').val())
        curTable=[]
        curSize=0
    }
}

function addClassToView(ind){
    console.log(ind)
    AddClass(curTable[parseInt(ind)])
}

function validateSub(){
    str = $('#subject_input').val().toUpperCase()
    len = $('#sub_drp li:contains("'+str+'")').length
    if(len!=1){
        stat_sub = false
        $('#subject_input').removeClass("valid")
        $('#subject_input').addClass("invalid")
    }
    else{
        stat_sub = true
        $('#subject_input').removeClass("invalid")
        $('#subject_input').addClass("valid")
    }
}

function validateCrse(){
    str = $('#course_input').val().toUpperCase()
    len = $('#crse_drp li:contains("'+str+'")').length
    if(len!=1){
        stat_crse = false
        $('#course_input').removeClass("valid")
        $('#course_input').addClass("invalid")
    }
    else{
        stat_crse = true
        $('#course_input').removeClass("invalid")
        $('#course_input').addClass("valid")
    }
}

$(document).ready(function() {
    
    //$('.modal-trigger').leanModal()
    $('#modal1').modal()
    //Init()
    $.ajax({
        type: "GET",
        url: "/subjects",
        async: false
    }).done(function(msg) {
        subjects_dict=msg.split('|')
        for (i in subjects_dict) {
            $('#sub_drp').append('<li><a onclick="setText(this,\'sub\')" ">'+subjects_dict[i]+'</a></li>')
        }
    })

    $('#subject_input').keyup(function(e) {
        if (e.which >= 65 && e.which <= 90 || e.which == 32 || e.which == 8) {
            str = this.value.toUpperCase()
            $('#sub_drp li').show()
            $('#sub_drp li:not(:contains("'+str+'"))').hide()
        }
    })

    $('#subject_input').focusout(function(){
        validateSub()
    })

    $('#course_input').focus(function(){
        subject = $('#subject_input').val().toUpperCase()
        $('#crse_drp').empty()
        $.ajax({
            type: "GET",
            url: "/course",
            data: {
                "subject": subject
            },
            async: false
        }).done(function(msg) {
            course_dict=msg.split('|')
            for (i in course_dict) {
                $('#crse_drp').append('<li><a onclick="setText(this,\'crse\')">'+course_dict[i]+'</a></li>')
            }
        })
    })
    $('#course_input').focusout(function(){
        validateCrse()
    })

    $('#course_input').keyup(function(e) {
        str = this.value.toUpperCase()
        $('#crse_drp li').show()
        $('#crse_drp li:not(:contains("'+str+'"))').hide()
    })
    f = false
    $('#modalTrigger').click(function(){
        $('#modal1').modal('open')
        if(!f){
            Init()
            f=true
        }
    })

})
