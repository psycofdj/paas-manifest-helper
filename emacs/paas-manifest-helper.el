(defcustom paas-manifest-helper-bin
  "paas-manifest-helper"
  "Path to paas-manifest-helper binary."
  :group 'paas-manifest-helper
  :type 'string
  :safe 'stringp)


;;;###autoload
(defun paas-manifest-helper-print-at-point()
  (interactive)
  (let* ((value (paas-manifest-helper-get-value-at-point)))
    (kill-new value)
    (message "%s" value))
  )


;;;###autoload
(defun paas-manifest-helper-open-at-point()
  (interactive)
  (let* ((value (paas-manifest-helper-get-value-at-point)))
    (if (file-exists-p value)
        (find-file value)
      (message "file not found: %s" value)))
  )

(defun paas-manifest-helper-get-value-at-point(&optional pline pcol)
  (let ((result "???")
        (line (if pline pline (number-to-string (line-number-at-pos))))
        (col  (if pcol  pcol  (number-to-string (current-column))))
        (outbuf (get-buffer-create "*paas-manifest-helper-result*")))

    (when (= 0 (call-process-region
                (point-min) (point-max) paas-manifest-helper-bin
                nil outbuf nil
                "--line" line "--col" col "--stdin" "--path" (buffer-file-name)))
      (with-current-buffer outbuf
        (setq result (replace-regexp-in-string "\n+" "" (buffer-string)))
        ))
    (kill-buffer outbuf)
    result
    )
  )

;; --------------------------------------------------------------------------- ;

;;;###autoload
(put 'paas-manifest-helper-bin 'safe-local-variable 'stringp)

(provide 'paas-manifest-helper)

;; Local Variables:
;; ispell-local-dictionary: "american"
;; End:
