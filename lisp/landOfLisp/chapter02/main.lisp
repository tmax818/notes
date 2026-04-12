;;;; It's convention to use four??




;;; create global variables

(defparameter *small* 1)
(defparameter *big* 100)

;;; define global functions


(defun guess-my-number ()
    (ash (+ *small* *big*) -1))