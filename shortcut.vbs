Set WshShell = WScript.CreateObject("WScript.Shell")
strDesktop = WshShell.SpecialFolders("Desktop")

If IsExitAFile(strDesktop & "\sdiary.lnk") Then
	CreateShortcut(strDesktop & "\sdiary.lnk")
End If

Function IsExitAFile(filespec)
	Dim fso
	Set fso=CreateObject("Scripting.FileSystemObject")        
	If fso.fileExists(filespec) Then         
	IsExitAFile=True        
	Else IsExitAFile=False        
	End If
End Function 

Sub CreateShortcut(filespec)
	set oShellLink = WshShell.CreateShortcut(filespec)
	oShellLink.TargetPath = "{{.}}"
	oShellLink.WindowStyle = 3 
	oShellLink.Save
End Sub
