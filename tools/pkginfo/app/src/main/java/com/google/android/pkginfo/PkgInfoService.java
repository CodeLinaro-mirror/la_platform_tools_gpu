package com.google.android.pkginfo;

import android.app.IntentService;
import android.content.Intent;
import android.content.pm.ActivityInfo;
import android.content.pm.ApplicationInfo;
import android.content.pm.PackageInfo;
import android.content.pm.PackageManager;
import android.graphics.Bitmap;
import android.graphics.drawable.BitmapDrawable;
import android.graphics.drawable.Drawable;
import android.net.LocalServerSocket;
import android.net.LocalSocket;
import android.util.Base64;
import android.util.Log;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.lang.reflect.Field;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * An {@link IntentService} subclass for providing installed package information to Android Studio.
 * <p/>
 * When the service is sent the {@link #ACTION_SEND_PKG_INFO} action, the service will begin
 * listening on the supplied local-abstract socket provided in the {@link #EXTRA_SOCKET_NAME} extra,
 * or if the extra is absent, {@link #DEFAULT_SOCKET_NAME}. When an incoming connection to this
 * socket is made, the service will send the installed package information on the accepted
 * connection, then close the accepted connection and the listening socket.
 */
public class PkgInfoService extends IntentService {
    private static final String TAG = "pkginfo";

    /**
     * Action used to start waiting for an incoming connection on the local-abstract port
     * {@link #EXTRA_SOCKET_NAME}. When a connection is made, the package information is send to the
     * connected socket, the socket is closed and the service stops listening on
     * {@link #EXTRA_SOCKET_NAME}.
     */
    private static final String ACTION_SEND_PKG_INFO = "com.google.android.pkginfo.action.SEND_PKG_INFO";

    /**
     * Optional parameter for {@link #ACTION_SEND_PKG_INFO} that changes the local-abstract port
     * used to listen for incoming connections. The default value is {@link #DEFAULT_SOCKET_NAME}.
     */
    private static final String EXTRA_SOCKET_NAME = "com.google.android.pkginfo.extra.SOCKET_NAME";

    /**
     * The default socket name when {@link #EXTRA_SOCKET_NAME} is not provided.
     */
    private static final String DEFAULT_SOCKET_NAME = "pkginfo";

    public PkgInfoService() {
        super("PkgInfoService");
    }

    @Override
    protected void onHandleIntent(Intent intent) {
        if (intent != null) {
            final String action = intent.getAction();
            if (ACTION_SEND_PKG_INFO.equals(action)) {
                String socketName = intent.getStringExtra(EXTRA_SOCKET_NAME);
                if (socketName == null) {
                    socketName = DEFAULT_SOCKET_NAME;
                }
                handleSendPackageInfo(socketName);
            }
        }
    }

    /**
     * Handler for the {@link #ACTION_SEND_PKG_INFO} intent.
     */
    private void handleSendPackageInfo(String socketName) {
        PackageManager pkgMgr = getPackageManager();
        List<PackageInfo> packages = pkgMgr.getInstalledPackages(PackageManager.GET_ACTIVITIES);

        // The ApplicationInfo.primaryCpuAbi field is hidden. Use reflection to get at it.
        Field primaryCpuAbiField = null;
        try {
            primaryCpuAbiField = ApplicationInfo.class.getField("primaryCpuAbi");
        } catch (NoSuchFieldException e) {
            Log.w(TAG, "Unable to find 'primaryCpuAbi' ApplicationInfo hidden field");
        }

        try {
            JSONArray packagesJson = new JSONArray();
            IconStore icons = new IconStore();

            for (PackageInfo packageInfo : packages) {
                Intent launchIntent = pkgMgr.getLaunchIntentForPackage(packageInfo.packageName);
                ActivityInfo launchActivityInfo = null;
                if (launchIntent != null) {
                    launchActivityInfo = launchIntent.resolveActivityInfo(pkgMgr, 0);
                }

                JSONArray activitiesJson = new JSONArray();
                if (packageInfo.activities != null) {
                    for (ActivityInfo activityInfo : packageInfo.activities) {
                        int iconIndex = -1;
                        if (activityInfo.icon > 0) {
                            iconIndex = icons.add(activityInfo.loadIcon(pkgMgr));
                        }

                        boolean isLaunchActivity = launchActivityInfo != null &&
                                activityInfo.name.equals(launchActivityInfo.name);

                        JSONObject activityJson = new JSONObject();
                        activityJson.put("Name", activityInfo.name);
                        activityJson.put("Icon", iconIndex);
                        activityJson.put("IsLaunch", isLaunchActivity);
                        activitiesJson.put(activityJson);
                    }
                }

                int iconIndex = -1;
                String primaryCpuAbi = null;
                ApplicationInfo applicationInfo = packageInfo.applicationInfo;
                if (applicationInfo != null) {
                    if (applicationInfo.icon > 0) {
                        iconIndex = icons.add(applicationInfo.loadIcon(pkgMgr));
                    }
                    if (primaryCpuAbiField != null) {
                        try {
                            primaryCpuAbi = (String) primaryCpuAbiField.get(applicationInfo);
                        } catch (Exception e) {
                            Log.w(TAG, "Exception thrown accessing 'primaryCpuAbi': " + e.getMessage());
                        }
                    }
                }

                JSONObject packageJson = new JSONObject();
                packageJson.put("Name", packageInfo.packageName);
                packageJson.put("Icon", iconIndex);
                if (primaryCpuAbi != null) {
                    packageJson.put("ABI", primaryCpuAbi);
                }
                packageJson.put("Activities", activitiesJson);
                packagesJson.put(packageJson);
            }

            JSONObject root = new JSONObject();
            root.put("Packages", packagesJson);
            root.put("Icons", icons.json());

            write(socketName, root.toString());
        } catch (JSONException e) {
            Log.e(TAG, e.toString());
        } catch (IOException e) {
            Log.e(TAG, e.toString());
        }
    }

    /**
     * IconStore stores all {@link Drawable}s as PNG, base-64 encoded images.
     * Duplicates are only stored once.
     */
    private class IconStore {
        private final Map<String, Integer> mMap = new HashMap();
        private final JSONArray mJson = new JSONArray();

        /**
         * add adds the specified drawable to the store.
         *
         * @return The index of the image stored in the {@link JSONArray} returned by {@link #json}.
         */
        public int add(Drawable drawable) {
            if (drawable == null || !(drawable instanceof BitmapDrawable)) {
                return -1;
            }
            Bitmap bitmap = ((BitmapDrawable) drawable).getBitmap();
            ByteArrayOutputStream stream = new ByteArrayOutputStream();
            bitmap.compress(Bitmap.CompressFormat.PNG, 100, stream);
            byte[] pngBytes = stream.toByteArray();
            String pngBase64 = Base64.encodeToString(pngBytes, Base64.NO_WRAP);
            if (!mMap.containsKey(pngBase64)) {
                int index = mJson.length();
                mMap.put(pngBase64, index);
                mJson.put(pngBase64);
                return index;
            } else {
                return mMap.get(pngBase64);
            }
        }

        /**
         * @return The {@link JSONArray} object holding all the base-64, PNG encoded images.
         */
        public JSONArray json() {
            return mJson;
        }
    }

    /**
     * write waits for the incoming connection to the specified local-abstract socket. When a
     * connection is made, data is written to the accepted connection, then both sockets are closed.
     *
     * @param socketName The name of the local-abstract socket to listen on.
     * @param data The data to send to the first accepted socket.
     */
    private void write(String socketName, String data) throws IOException {
        LocalServerSocket server = new LocalServerSocket(socketName);
        try {
            LocalSocket socket = server.accept();
            try {
                socket.getOutputStream().write(data.getBytes());
            } finally {
                socket.close();
            }
        } finally {
            server.close();
        }
    }
}
