    $(document).ready(function(){
       DoUpdate(); 
    });
    function DoUpdate(){
        $.ajax({url: "vals", success: function(data){parsevalues(data)}
    });
    }
function SetValue(valname, val){
    console.log("Setting" + valname + " to " + val);
    turl = "vals/" + valname + "?val=" + val
    $.ajax({url: turl, success: function(data){setresp(data)}
});
}
function setresp(data){
    console.log(data);
}
function rotate(divname,deg) {
    var rotated = false;

   
        var div = document.getElementById(divname);
            
    
        div.style.webkitTransform = 'rotate('+deg+'deg)'; 
        div.style.mozTransform    = 'rotate('+deg+'deg)'; 
        div.style.msTransform     = 'rotate('+deg+'deg)'; 
        div.style.oTransform      = 'rotate('+deg+'deg)'; 
        div.style.transform       = 'rotate('+deg+'deg)'; 
    
        rotated = !rotated;
       
  }
function parsevalues(data){
    jpl = JSON.parse(data);
    //alert(Math.round(jpl['APHEADING']));
    $("#NAV1HDEF").html(jpl['NAV1HDEF']);
    $("#NAV1VDEF").html(jpl['NAV1VDEF']);
    $("#NAV1DMED").html(jpl['NAV1DMED']);
    
    $("#NAV1CURRENT").html((jpl['NAV1CurrentFrequency']/100));
    $("#NAV1NEXT").html((jpl['NAV1StandbyFrequency']/100));
    $("#NAV2CURRENT").html((jpl['NAV2CurrentFrequency']/100));
    $("#NAV2NEXT").html((jpl['NAV2StandbyFrequency']/100));

    //
    $("#APHEADINGVAL").html(Math.round(jpl['APHEADING']));
        rotate("GAUGEPOINTER",- jpl['ROLL']);
        rotate("GAUGEBACK",- jpl['ROLL']);
       
     $("#GAUGEBACK").css({ top: (jpl['PITCH'])*2 +'px' });
    setTimeout(function(){ DoUpdate(); }, 100);
}
function flipnav1(){
    var oldcurrent = $("#NAV1CURRENT").html();
    
    SetValue("NAV1CurrentFrequency", ($("#NAV1NEXT").html()*100))
    SetValue("NAV1StandbyFrequency", (oldcurrent*100))

}
function flipnav2(){
    var oldcurrent = $("#NAV2CURRENT").html();
    
    SetValue("NAV2CurrentFrequency", ($("#NAV2NEXT").html()*100))
    SetValue("NAV2StandbyFrequency", (oldcurrent*100))

}
function hdgleft1(){
    newval = parseInt($("#APHEADINGVAL").html()) - 1;
    if (newval < 0){
        newval = newval + 360
    }
    //alert(newval);
    SetValue("APHEADING", newval);
}
function hdgleft10(){
    newval = parseInt($("#APHEADINGVAL").html()) - 10;
    if (newval < 0){
        newval = newval + 360
    }
    //alert(newval);
    SetValue("APHEADING", newval);
}
function hdgright1(){
    newval = parseInt($("#APHEADINGVAL").html()) + 1;
    if (newval > 360){ newval = newval - 360}
    //alert(newval);
    SetValue("APHEADING", newval);
}
function hdgright10(){
    newval = parseInt($("#APHEADINGVAL").html()) + 10;
    if (newval > 360){ newval = newval - 360}
    //alert(newval);
    SetValue("APHEADING", newval);
}