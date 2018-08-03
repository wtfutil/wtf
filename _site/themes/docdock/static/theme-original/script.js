// Scrollbar Width function
function getScrollBarWidth() {
    var inner = document.createElement('p');
    inner.style.width = "100%";
    inner.style.height = "200px";

    var outer = document.createElement('div');
    outer.style.position = "absolute";
    outer.style.top = "0px";
    outer.style.left = "0px";
    outer.style.visibility = "hidden";
    outer.style.width = "200px";
    outer.style.height = "150px";
    outer.style.overflow = "hidden";
    outer.appendChild(inner);

    document.body.appendChild(outer);
    var w1 = inner.offsetWidth;
    outer.style.overflow = 'scroll';
    var w2 = inner.offsetWidth;
    if (w1 == w2) w2 = outer.clientWidth;

    document.body.removeChild(outer);

    return (w1 - w2);
};

// for the window resize
$(window).resize(function() {});

// debouncing function from John Hann
// http://unscriptable.com/index.php/2009/03/20/debouncing-javascript-methods/
(function($, sr) {

    var debounce = function(func, threshold, execAsap) {
        var timeout;

        return function debounced() {
            var obj = this, args = arguments;

            function delayed() {
                if (!execAsap)
                    func.apply(obj, args);
                timeout = null;
            };

            if (timeout)
                clearTimeout(timeout);
            else if (execAsap)
                func.apply(obj, args);

            timeout = setTimeout(delayed, threshold || 100);
        };
    }
    // smartresize
    jQuery.fn[sr] = function(fn) { return fn ? this.bind('resize', debounce(fn)) : this.trigger(sr); };

})(jQuery, 'smartresize');

jQuery(document).ready(function() {

    var sidebarStatus = searchStatus = 'open';

    jQuery('#overlay').on('click', function() {
        jQuery(document.body).toggleClass('sidebar-hidden');
        sidebarStatus = (jQuery(document.body).hasClass('sidebar-hidden') ? 'closed' : 'open');

        return false;
    });

    jQuery('[data-sidebar-toggle]').on('click', function() {
        jQuery(document.body).toggleClass('sidebar-hidden');
        sidebarStatus = (jQuery(document.body).hasClass('sidebar-hidden') ? 'closed' : 'open');

        return false;
    });

    jQuery('[data-search-toggle]').on('click', function() {
        if (sidebarStatus == 'closed') {
            jQuery('[data-sidebar-toggle]').trigger('click');
            jQuery(document.body).removeClass('searchbox-hidden');
            searchStatus = 'open';

            return false;
        }

        jQuery(document.body).toggleClass('searchbox-hidden');
        searchStatus = (jQuery(document.body).hasClass('searchbox-hidden') ? 'closed' : 'open');

        return false;
    });

    var touchsupport = ('ontouchstart' in window) || (navigator.maxTouchPoints > 0) || (navigator.msMaxTouchPoints > 0)
    if (!touchsupport){ // browser doesn't support touch
        $('#toc-menu').hover(function() {
            $('.progress').stop(true, false, true).fadeToggle(100);
        });

        $('.progress').hover(function() {
            $('.progress').stop(true, false, true).fadeToggle(100);
        });
    }
    if (touchsupport){ // browser does support touch
        $('#toc-menu').click(function() {
            $('.progress').stop(true, false, true).fadeToggle(100);
        });
        $('.progress').click(function() {
            $('.progress').stop(true, false, true).fadeToggle(100);
        });
    }
});

jQuery(window).on('load', function() {

    function adjustForScrollbar() {
        if ((parseInt(jQuery('#body-inner').height()) + 83) >= jQuery('#body').height()) {
            jQuery('.nav.nav-next').css({ 'margin-right': getScrollBarWidth() });
        } else {
            jQuery('.nav.nav-next').css({ 'margin-right': 0 });
        }
    }

    // adjust sidebar for scrollbar
    adjustForScrollbar();

    jQuery(window).smartresize(function() {
        adjustForScrollbar();
    });


    
    
});
