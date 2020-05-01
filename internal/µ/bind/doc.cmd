@Echo Off

echo // go doc put to good use	 >doc.md
echo.				>>doc.md
echo ------------------------------------------------------------------------------->>doc.md

echo ## go doc    bind.Ings	>>doc.md
        go doc    bind.Ings	>>doc.md

echo.				>>doc.md
echo ------------------------------------------------------------------------------->>doc.md

echo ## go doc -u bind.Ings	>>doc.md
        go doc -u bind.Ings	>>doc.md

echo.				>>doc.md
echo ------------------------------------------------------------------------------->>doc.md

echo ## go doc -all		>>doc.md
        go doc -all		>>doc.md

echo.				>>doc.md
echo ------------------------------------------------------------------------------->>doc.md
