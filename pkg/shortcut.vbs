Set WshShell = WScript.CreateObject("WScript.Shell")
strDesktop = WshShell.SpecialFolders("Desktop")

If IsExitAFile(strDesktop & "\secret-diary.lnk") Then
	rem "exists"
Else
	CreateShortcut(strDesktop & "\secret-diary.lnk")
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
	Set WshShell = WScript.CreateObject("WScript.Shell")
	Set oShellLink = WshShell.CreateShortcut(filespec)
	oShellLink.TargetPath = Wscript.Arguments.Named("target")
	oShellLink.WindowStyle = 3 
	oShellLink.Save
End Sub
