package com.vladimirdobro.deploymanager

import android.app.Activity
import android.content.Intent
import android.net.Uri
import android.os.Bundle
import android.webkit.ValueCallback
import android.webkit.WebChromeClient
import android.webkit.WebResourceRequest
import android.webkit.WebView
import android.webkit.WebViewClient
import androidx.activity.result.contract.ActivityResultContracts
import androidx.appcompat.app.AppCompatActivity
import androidx.core.content.FileProvider
import backend.Backend
import java.io.File

class MainActivity : AppCompatActivity() {
    private var fileUploadCallback: ValueCallback<Array<Uri>>? = null

    private val fileUploadActivityResultLauncher = registerForActivityResult(ActivityResultContracts.StartActivityForResult()) { result ->
        if (result.resultCode == Activity.RESULT_OK) {
            val results = result.data?.let { WebChromeClient.FileChooserParams.parseResult(result.resultCode, it) }
            fileUploadCallback?.onReceiveValue(results)
        } else {
            fileUploadCallback?.onReceiveValue(null)
        }
        fileUploadCallback = null
    }

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        val webView: WebView = findViewById(R.id.webview)
        webView.settings.javaScriptEnabled = true
        // Upload file from the WebView
        webView.webChromeClient = object : WebChromeClient() {
            override fun onShowFileChooser(
                webView: WebView?,
                filePathCallback: ValueCallback<Array<Uri>>?,
                fileChooserParams: FileChooserParams?
            ): Boolean {
                fileUploadCallback = filePathCallback ?: return false
                val intent = fileChooserParams?.createIntent()
                fileUploadActivityResultLauncher.launch(intent)
                return true
            }
        }
        // Download file in default browser
        webView.webViewClient = object : WebViewClient() {
            override fun shouldOverrideUrlLoading(view: WebView?, request: WebResourceRequest?): Boolean {
                val openURL = Intent(Intent.ACTION_VIEW)
                val url = request?.url.toString()
                if (url.contains("deploy")) {
                    openURL.data = Uri.parse(url)
                    startActivity(openURL)
                }
                else {
                    view?.loadUrl(url)
                }
                return true
            }
        }

        Backend.run("0.0.0.0:2082")
        webView.loadUrl("http://127.0.0.1:2082/www")
//        val openDIR = Intent(Intent.ACTION_GET_CONTENT)
//        val dir = Uri.parse("/storage/emulated/0/")
//        openDIR.setDataAndType(dir, "*/*")
//        startActivity(Intent.createChooser(openDIR, "Open folder"))
//        val file = File("/storage/emulated/0/Download/test.txt")
//        if(file.exists()) {
//            val uri = FileProvider.getUriForFile(this, BuildConfig.APPLICATION_ID + ".provider", file)
//            val intent = Intent(Intent.ACTION_SEND)
//            intent.addFlags(Intent.FLAG_GRANT_READ_URI_PERMISSION or Intent.FLAG_GRANT_WRITE_URI_PERMISSION)
//            intent.type = "*/*"
//            intent.putExtra(Intent.EXTRA_STREAM, uri)
//            intent.flags = Intent.FLAG_ACTIVITY_NEW_TASK;
//            startActivity(intent)
//        }
    }
}
