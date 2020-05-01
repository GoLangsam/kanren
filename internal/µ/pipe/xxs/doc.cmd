@Echo Off

echo // go doc put to good use	 >doc.md
echo.				>>doc.md
echo ------------------------------------------------------------------------------->>doc.md

echo ## go doc .  StreamOfStates	>>doc.md
	go doc github.com/GoLangsam/kanren/internal/æ/pipe/xxs StreamOfStates	>>doc.md

echo.				>>doc.md
echo ------------------------------------------------------------------------------->>doc.md

echo ## go doc -all		>>doc.md
	go doc -all github.com/GoLangsam/kanren/internal/æ/pipe/xxs		>>doc.md

echo.				>>doc.md
echo ------------------------------------------------------------------------------->>doc.md
